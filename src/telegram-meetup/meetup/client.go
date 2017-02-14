/*
https://api.meetup.com/Milano-Chatbots-Meetup/events?photo-host=public&page=20&key=564a392576101d105434415277167e

*/

package meetup

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	. "telegram-meetup/types"
	"time"

	"errors"

	"bytes"
	"crypto/md5"
	"strconv"

	"golang.org/x/net/context"
)

type HttpDo func(context.Context, *http.Request) (*http.Response, error)

type Client struct {
	env            Env
	c              HttpDo
	apiEndpoint    string
	key            string
	oauth2AuthEP   string
	oauth2AccessEP string
	clientID       string
	clientSecret   string
}

func string2url(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func NewClient(env Env, c HttpDo) *Client {
	var token = env.Get("MEETUP_KEY")
	if len(token) < 5 {
		panic("undefined or invalid Meetup key: " + token)
	}
	var endpoint = env.Get("MEETUP_ENDPOINT")
	if len(endpoint) == 0 {
		endpoint = "https://api.meetup.com/"
	}

	var oauth2AuthEP = env.Get("MEETUP_OAUTH2_AUTH_EP")
	if len(oauth2AuthEP) == 0 {
		oauth2AuthEP = "https://secure.meetup.com/oauth2/authorize"
	}
	var oauth2AccessEP = env.Get("MEETUP_OAUTH2_ACCESS_EP")
	if len(oauth2AccessEP) == 0 {
		oauth2AccessEP = "https://secure.meetup.com/oauth2/access"
	}

	var clientID = env.Get("MEETUP_CLIENT_ID")
	var clientSecret = env.Get("MEETUP_CLIENT_SECRET")
	return &Client{c: c,
		env:            env,
		apiEndpoint:    endpoint,
		key:            token,
		oauth2AuthEP:   oauth2AuthEP,
		oauth2AccessEP: oauth2AccessEP,
		clientID:       clientID,
		clientSecret:   clientSecret,
	}
}
func (c *Client) get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.c(ctx, req)
}
func (c *Client) postForm(ctx context.Context, url string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	return c.c(ctx, req)
}

func (c *Client) GetNextMeetup(ctx context.Context, group MeetupGroupID, ref time.Time) (*Meetup, error) {

	type entry struct {
		Name  string
		Time  int64
		Venue struct {
			Address_1 string
			City      string
			Lon       float64
			Lat       float64
		}
		ID      string
		Updated int64
	}
	var results = make([]entry, 0, 10)
	var url = fmt.Sprintf("%s/%s/events?photo-host=public&page=20&key=%s",
		c.apiEndpoint, group, c.key,
	)

	resp, err := c.get(ctx, url)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBody, &results)
	if err != nil {
		return nil, err
	}

	//return nil, nil

	if len(results) > 0 {
		r0 := results[0]
		when := time.Unix(r0.Time/1000, 0)
		var mockWhen = c.env.Get("MEETUP_MOCK_WHEN")
		if len(mockWhen) > 0 {
			var intVal, _ = strconv.Atoi(mockWhen)
			when = time.Unix(int64(intVal), 0)
		}

		l, err := time.LoadLocation("Europe/Rome")
		if err == nil {
			when = when.In(l)
		}

		first := &Meetup{
			Title: r0.Name,
			Where: r0.Venue.Address_1 + ", " + r0.Venue.City,
			When:  when,
			ETag:  "not yet",
			Map:   &Location{Lat: r0.Venue.Lat, Lng: r0.Venue.Lon},
		}
		first.ETag = md5hex(
			fmt.Sprintf("%s %s %v %v", first.Title, first.Where, first.When, first.Map),
		)
		if first.When.Before(ref) {
			first.ETag = "x" + first.ETag
		}

		return first, nil
	}
	return nil, nil
}
func md5hex(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *Client) MakeAuthorizeURL(redirectUri string, state string) (URL, error) {
	var q = url.Values{}
	q.Set("redirect_uri", redirectUri)
	q.Set("state", state)
	q.Set("response_type", "code")
	q.Set("client_id", c.clientID)
	u := c.oauth2AuthEP + "?" + q.Encode()
	return URL(u), nil
}
func (c *Client) RefreshAuthTokens(context.Context, MeetupAuthTokens) (MeetupAuthTokens, error) {
	panic("not implemented")
}

func (c *Client) GetGroupList(ctx context.Context, t MeetupAuthTokens) ([]MeetupGroup, error) {
	_, access, _ := unpackTokens(t)
	type entry struct {
		Name    string `json:"name"`
		Urlname string `json:"urlname"`
		ID      int    `json:"id"`
	}

	var results = make([]entry, 0, 10)
	var url = fmt.Sprintf("%s/self/groups?access_token=%s",
		c.apiEndpoint, access,
	)
	resp, err := c.get(ctx, url)
	log.Println(resp)
	respBody, err := ioutil.ReadAll(resp.Body)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		if resp.StatusCode == 401 {
			return nil, ErrMeetup401
		} else {
			return nil, errors.New(resp.Status)
		}
	}
	err = json.Unmarshal(respBody, &results)
	log.Println(err)
	if err != nil {
		return nil, err
	}
	rv := make([]MeetupGroup, 0, len(results))
	for _, v := range results {
		rv = append(rv, MeetupGroup{v.Name, v.Urlname, v.ID})
	}
	return rv, nil
}

type oauth2AccessResult struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type"`
	ExpiresIn    time.Duration `json:"expires_in"`
	RefreshToken string        `json:"refresh_token"`
}
type oauth2AccessError struct {
	Error string `json:"error"`
}

func (c *Client) ProcessAuthResult(ctx context.Context, arg url.Values, redirectUri string) (MeetupAuthTokens, error) {
	var args = struct{ Code, Error string }{
		arg.Get("code"), arg.Get("error"),
	}

	if len(args.Error) > 0 {
		return nil, errors.New(args.Error)
	}

	var payload = url.Values{
		"client_id":     []string{c.clientID},
		"client_secret": []string{c.clientSecret},
		"grant_type":    []string{"authorization_code"},
		"redirect_uri":  []string{redirectUri},
		"code":          []string{args.Code},
	}

	var url = c.oauth2AccessEP

	resp, err := c.postForm(ctx, url, payload)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		var result oauth2AccessError
		err := json.Unmarshal(respBody, &result)
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New(result.Error)
		}
	}
	var result oauth2AccessResult
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}
	return packTokens(time.Now().Add(result.ExpiresIn*time.Second), result.AccessToken, result.RefreshToken), nil
}

func packTokens(exp time.Time, access, refresh string) MeetupAuthTokens {
	return bytes.Join([][]byte{
		[]byte(strconv.Itoa(int(exp.Unix()))),
		[]byte(access),
		[]byte(refresh),
	}, []byte(";"))
}

func DebugTokens(t MeetupAuthTokens) (exp time.Time, access, refresh string) {
	return unpackTokens(t)
}

func unpackTokens(t MeetupAuthTokens) (exp time.Time, access, refresh string) {
	var chunks = bytes.Split(t, []byte(";"))
	if len(chunks) == 3 {
		i, _ := strconv.Atoi(string(chunks[0]))
		exp = time.Unix(int64(i), 0)
		access = string(chunks[1])
		refresh = string(chunks[1])
	}
	_ = chunks
	return
}
