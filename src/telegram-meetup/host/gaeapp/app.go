package gaeapp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"telegram-meetup/env"
	"telegram-meetup/meetup"
	"telegram-meetup/telegram"
	"telegram-meetup/thebot"
	. "telegram-meetup/types"
	"time"

	"github.com/pborman/uuid"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const InitHandler = "/init"

const TelegramWebhook = "/telegram-webhook"
const FacebookWebhook = "/facebook-webhook"
const DatastoreVolatileKind = "volatile"
const MeetupOauth2Start = "/oauth/start"
const MeetupOauth2Return = "/oauth/return"

const CronHandler = "/cron-handler"

const SendMessageHandler = "/send-message"

const EnvMyBaseURL = "MY_BASE_URL"

func discoverMyBaseURL(r *http.Request) string {
	proto := nvl(r.Header.Get("X-Forwarded-Proto"), "https")
	hostname := nvl(r.Header.Get("X-Appengine-Server-Name"), r.Header.Get("Host"))
	log.Println(proto)
	log.Println(hostname)
	return proto + "://" + hostname
}

func makeClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}

func httpDo(ctx context.Context, r *http.Request) (*http.Response, error) {
	return urlfetch.Client(ctx).Do(r)
}

func init() {
	if len(os.Getenv("RUN_WITH_DEVAPPSERVER")) > 0 {
		env.DevMode = true
	}
	log.Printf("%v DevMode=%v", os.Environ(), env.DevMode)
	var env = env.LoadEnv()

	http.HandleFunc(TelegramWebhook,
		func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(461)
				fmt.Fprint(w, err)
			} else {
				ctx := appengine.NewContext(r)
				mq := NewMessageQueue()
				mq.HandleWebhookMessage(ctx, &WebhookMessage{"telegram", body})
				/*
					bot, _ := telegram.NewBot(env, mq, makeClient)

					err := bot.HandleWebhookPayload(ctx, body)
					if err != nil {
						w.WriteHeader(462)
						fmt.Fprint(w, err)
					}
				*/
			}

		},
	)

	http.HandleFunc(InitHandler,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(200)
			ctx := appengine.NewContext(r)
			proto := nvl(r.Header.Get("X-Forwarded-Proto"), "https")
			hostname := nvl(r.Header.Get("X-Appengine-Server-Name"), r.Header.Get("Host"))
			log.Println(proto)
			log.Println(hostname)
			url := proto + "://" + hostname
			bot, _ := telegram.NewBot(env, nil, makeClient)

			username, err := bot.SetWebhook(ctx, url+TelegramWebhook)

			botUrl := "http://telegram.me/" + username + "?start=qrcode"

			log.Println(err)
			/*out, _ := json.Marshal(&r.Header)
			fmt.Fprint(w, string(out))*/
			png, err := qrcode.Encode(botUrl, qrcode.Medium, 800)
			fmt.Fprint(w, url, "\n")
			fmt.Fprint(w, botUrl)
			var mq = NewMessageQueue()
			err = mq.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: env.Get("ADMIN_CHAT_ID"),
				Message:     "Callback: " + url + TelegramWebhook,
				Attachment:  &File{username + ".png", png},
			})
			fmt.Fprint(w, err)
			glog.Debugf(ctx, "init: err %v", err)

		},
	)

	http.HandleFunc("/qrcode.png",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(200)
			ctx := appengine.NewContext(r)

			bot, _ := telegram.NewBot(env, nil, makeClient)
			api, _ := bot.MakeApi(ctx)
			me, _ := api.GetMe()

			botUrl := "http://telegram.me/" + me.UserName + "?start=qrcode"

			png, err := qrcode.Encode(botUrl, qrcode.Medium, 500)
			w.Write(png)
			glog.Debugf(ctx, "init: err %v", err)

		},
	)
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(200)
			ctx := appengine.NewContext(r)

			bot, _ := telegram.NewBot(env, nil, makeClient)
			api, _ := bot.MakeApi(ctx)
			me, _ := api.GetMe()
			botUrl := "http://t.me/" + me.UserName + "?start=qrcode"

			fmt.Fprintf(w, `  
			<body style="color: black">
			<center><h1><img src="%s"><a style="color: black" href='%s'>@%s</a></h1><img width=500 height=500 src="/qrcode.png"></center>
			</body>
			`, telegram.TelegramIconData, botUrl, me.UserName)

		},
	)

	http.HandleFunc(CronHandler,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := appengine.NewContext(r)
			_, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(471)
				glog.Debugf(ctx, "err: %v", err)
				fmt.Fprint(w, err)
				return

			}
			var volatiles = datastoreVolatile{ctx}
			err = volatiles.deleteExpiredEntries()
			if err != nil {
				w.WriteHeader(472)
				glog.Debugf(ctx, "err: %v", err)
				fmt.Fprint(w, err)
				return
			}

			var mq = NewMessageQueue()
			mq.HandleJobRequest(ctx, &JobRequest{
				Command:   []string{"check_new_notify"},
				Scheduled: time.Now(),
				Droppable: time.Now().Add(1 * time.Hour),
				Cookie:    nil,
			})
			if err != nil {
				w.WriteHeader(463)
				panic(err)
			}
			w.WriteHeader(200)
			fmt.Fprint(w, "ok")

		})

	http.HandleFunc(MeetupOauth2Start,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(461)
				fmt.Fprint(w, err)
			} else {
				q := r.URL.Query()
				key := q.Get("key")
				ctx := appengine.NewContext(r)
				var volatiles = datastoreVolatile{ctx}
				var mc, _ = NewMeetupCollaborators(env)

				var state, _ = volatiles.getBinaryData(key)
				myBaseURL := env.Get(EnvMyBaseURL)
				var authURL, _ = mc.MakeAuthorizeURL(myBaseURL+MeetupOauth2Return, key)

				w.Header().Add("Location", string(authURL))
				w.WriteHeader(301)

				fmt.Fprint(w, state, "\n")
				fmt.Fprint(w, authURL, "\n")
			}

		})
	http.HandleFunc(MeetupOauth2Return,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(461)
				fmt.Fprint(w, err)
			} else {

				q := r.URL.Query()
				key := q.Get("state")
				code := q.Get("code")
				ctx := appengine.NewContext(r)
				var volatiles = datastoreVolatile{ctx}
				var mc, ma = NewMeetupCollaborators(env)
				var us = NewUserStore()
				var state, _ = volatiles.getBinaryData(key)
				var mq = NewMessageQueue()

				fmt.Fprint(w, state, "\n")
				fmt.Fprint(w, code, "\n")
				myBaseURL := env.Get(EnvMyBaseURL)

				log.Println(state, code, myBaseURL)

				redirectUri := myBaseURL + MeetupOauth2Return
				tokens, err := mc.ProcessAuthResult(ctx, q, redirectUri)

				_ = tokens
				_ = mq
				_ = mc
				_ = ma
				_ = us

				if err != nil {
					glog.Debugf(ctx, "ProcessAuthResult: %v", err)
				} else {
					glog.Debugf(ctx, "Tokens: %v", string(tokens))
					log.Println(meetup.DebugTokens(tokens))
					var engine = thebot.NewBot(env, us, mq, mc, ma)
					engine.HandleMeetupAuthCompletion(
						ctx,
						&MeetupAuthCompletion{
							State:  state,
							Tokens: tokens,
						},
					)
				}

			}

		})

	http.HandleFunc(FacebookWebhook,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(499)
				fmt.Fprint(w, err)
			} else {
				ctx := appengine.NewContext(r)
				q := r.URL.Query()
				key := q.Get("hub.challenge")
				w.WriteHeader(200)
				fmt.Fprint(w, key)
				glog.Debugf(ctx, "%v", string(body))
			}
		})

	http.HandleFunc(SendMessageHandler,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(499)
				fmt.Fprint(w, err)
			} else {
				ctx := appengine.NewContext(r)
				q := r.URL.Query()
				key := q.Get("key")
				if len(key) > 0 && key == env.Get("AUTH_KEY") {
					recipient := q.Get("recipient")
					var mq = NewMessageQueue()
					if recipient == "*" {
						mq.HandleJobRequest(ctx, &JobRequest{
							Command:   []string{"broadcast_current_event", string(body)},
							Cookie:    nil,
							Scheduled: time.Now(),
							Droppable: time.Now().Add(60 * time.Second),
						})
					} else {
						mq.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
							RecipientID: recipient,
							Message:     string(body),
						})
					}
				} else {
					w.WriteHeader(403)
					fmt.Fprint(w, "invalid key")
					//glog.Debugf(ctx, "%v", string(body))
				}
			}
		})

	SetupMessageQueueHandler(env)

}

/* ################################################################################ */
// handler factory

type handlerFactory struct {
}

/* ############################################################################################### */

/* ---------------------------------------------------------------------- */

type MyMeetupAuthorizer struct {
	meetup    *meetup.Client
	myBaseURL string
}

func (a *MyMeetupAuthorizer) MakeOOBAuthStartURL(ctx context.Context, state MeetupAuthState) (URL, error) {
	key := uuid.New()
	var volatiles = datastoreVolatile{ctx}
	err := volatiles.setBinaryData(key, state, 360*time.Second)
	if err != nil {
		return URL(""), err
	}
	q := url.Values{}
	q.Set("key", key)
	var u = a.myBaseURL + MeetupOauth2Start + "?" + q.Encode()
	return URL(u), nil
}

func (c *MyMeetupAuthorizer) RefreshAuthTokens(ctx context.Context, tokens MeetupAuthTokens) (MeetupAuthTokens, error) {

	return c.meetup.RefreshAuthTokens(ctx, tokens)
}

func NewMeetupCollaborators(env Env) (*meetup.Client, *MyMeetupAuthorizer) {
	myBaseURL := env.Get(EnvMyBaseURL)
	var mc = meetup.NewClient(env, httpDo)
	var ma = &MyMeetupAuthorizer{mc, myBaseURL}
	return mc, ma
}
