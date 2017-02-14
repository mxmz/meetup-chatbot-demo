package meetup

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"

	. "github.com/onsi/gomega"
)

type logMap map[string]string

type osEnv struct {
}

func (*osEnv) Get(k string) string {
	return os.Getenv(k)
}

func defaultHttpDo(ctx context.Context, r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r.WithContext(ctx))
}

func startTestServer(mockRespBody []byte, testLog logMap) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testLog["url"] = r.URL.String()
		body, _ := ioutil.ReadAll(r.Body)
		testLog["body"] = string(body)
		w.Write(mockRespBody)
	}))
	return ts
}

func Test_GetNextMeetup_SomeResults(t *testing.T) {
	RegisterTestingT(t)
	var testLog = make(logMap)
	var mockRespBody = []byte(responseOk)
	ts := startTestServer(mockRespBody, testLog)
	defer ts.Close()
	os.Setenv("MEETUP_ENDPOINT", ts.URL)
	os.Setenv("MEETUP_KEY", "T3STK3Y")
	var client = NewClient(&osEnv{}, defaultHttpDo)
	var meetup, err = client.GetNextMeetup(context.TODO(), "Milano-Chatbots-Meetup", time.Now())
	Expect(err).To(BeNil())
	Expect(meetup).NotTo(BeNil())
	Expect(meetup.Title).To(BeIdenticalTo("Bot, sviluppo e commercio"))
	Expect(meetup.Where).To(BeIdenticalTo("via Venini 42, Milano"))
	Expect(meetup.When.Unix()).To(BeEquivalentTo(time.Unix(1481824800, 0).Unix()))
	Expect(meetup.ETag).To(BeIdenticalTo("x078ac95bd66c8a4b1ec51bee5e38c778"))
	var expectedLog = logMap{
		"body": "",
		"url":  "/Milano-Chatbots-Meetup/events?photo-host=public&page=20&key=T3STK3Y",
	}
	Expect(testLog).To(BeEquivalentTo(expectedLog))
}

func Test_GetNextMeetup_0Results(t *testing.T) {
	RegisterTestingT(t)

	var testLog = make(logMap)
	var mockRespBody = []byte(`[]`)
	ts := startTestServer(mockRespBody, testLog)
	defer ts.Close()

	os.Setenv("MEETUP_ENDPOINT", ts.URL)
	os.Setenv("MEETUP_KEY", "T3STK3Y")

	var client = NewClient(&osEnv{}, defaultHttpDo)

	var meetup, err = client.GetNextMeetup(context.TODO(), "Milano-Chatbots-Meetup", time.Now())
	Expect(err).To(BeNil())
	Expect(meetup).To(BeNil())

	var expectedLog = logMap{
		"body": "",
		"url":  "/Milano-Chatbots-Meetup/events?photo-host=public&page=20&key=T3STK3Y",
	}
	Expect(testLog).To(BeEquivalentTo(expectedLog))

}

func Test_GetNextMeetup_InvalidResponse(t *testing.T) {
	RegisterTestingT(t)
	var testLog = make(logMap)
	var mockRespBody = []byte(`{}`)
	ts := startTestServer(mockRespBody, testLog)
	defer ts.Close()
	os.Setenv("MEETUP_ENDPOINT", ts.URL)
	os.Setenv("MEETUP_KEY", "T3STK3Y")
	var client = NewClient(&osEnv{}, defaultHttpDo)
	var meetup, err = client.GetNextMeetup(context.TODO(), "Milano-Chatbots-Meetup", time.Now())
	Expect(meetup).To(BeNil())
	Expect(err.Error()).To(MatchRegexp(`cannot unmarshal object`))
	var expectedLog = logMap{
		"body": "",
		"url":  "/Milano-Chatbots-Meetup/events?photo-host=public&page=20&key=T3STK3Y",
	}
	Expect(testLog).To(BeEquivalentTo(expectedLog))

}

func Test_GetNextMeetup_EmptyResponse(t *testing.T) {
	RegisterTestingT(t)
	var testLog = make(logMap)
	var mockRespBody = []byte(``)
	ts := startTestServer(mockRespBody, testLog)
	defer ts.Close()
	os.Setenv("MEETUP_ENDPOINT", ts.URL)
	os.Setenv("MEETUP_KEY", "T3STK3Y")
	var client = NewClient(&osEnv{}, defaultHttpDo)
	var meetup, err = client.GetNextMeetup(context.TODO(), "Milano-Chatbots-Meetup", time.Now())
	Expect(meetup).To(BeNil())
	Expect(err.Error()).To(MatchRegexp(`unexpected end of JSON input`))
	var expectedLog = logMap{
		"body": "",
		"url":  "/Milano-Chatbots-Meetup/events?photo-host=public&page=20&key=T3STK3Y",
	}
	Expect(testLog).To(BeEquivalentTo(expectedLog))

}

const responseOk = `
[
   {
      "name" : "Bot, sviluppo e commercio",
      "how_to_find_us" : "MM1 Pasteur o tram 1 fermata Venini/Sauli. Citofono Mikamai/LinkMe/Venini42.",
      "status" : "upcoming",
      "description" : "<p>Nel prossimo meetup di Milano Chatbots avremo una presentazione di Giorgio Robino, che con il <a href=\"https://medium.com/convcomp2016\">convcomp2016</a> di giugno ha dato il via a tutta una serie di meetup tra cui anche i nostri di Milano. Il suo talk e quello conclusivo di Claudio Comandini saranno incentrati sulla costruzione di bot. Tra di loro Simone Guzzetti ci racconterÃ  la sua esperienza imprenditoriale con <a href=\"http://www.bloovery.com/it/\">bloovery.com</a></p> <p><br/>I talk in programma sono:</p> <p><br/>• Keynote di Paolo Montrasio.</p> <p><br/>• \"naif - Ruby micro-framework to build dumb chat-machines\", di Giorgio Robino.</p> <p>• \"Un bot per vendere fiori\", di Simone Guzzetti.</p> <p>• \"Un bot multi channel con assi.st\", di Claudio Comandini, sessione di sviluppo live.</p> <p>\n\nDopo i talk ci saranno le solite due ore di networking con l'apericena a buffet autogestito. Mi raccomando, portate qualcosa da bere e da mangiare per renderle piÃ¹ piacevoli!</p> <p>Infine, chi volesse tenere un talk nei prossimi mesi mi contatti al piÃ¹ presto per bloccare uno slot nei meetup di gennaio e febbraio.</p> <p>Al 15 dicembre!</p> <p>Paolo</p> <p>\n\n\nPS: chi parteciperÃ  al Codemotion di Milano non dimentichi di venire al nostro evento Chatbot di venerdÃ¬ 25 novembre alle 18:30 (<a href=\"https://www.meetup.com/Milano-Chatbots-Meetup/messages/boards/thread/50395870\">programma qui</a>) e al talk di Paolo Montrasio di sabato 26 alle 16:20, in cui sarÃ  dimostrata la costruzione di un bot Telegram su AWS Lambda.</p> ",
      "venue" : {
         "name" : "Mikamai e LinkMe",
         "city" : "Milano",
         "country" : "it",
         "localized_country_name" : "Italia",
         "lat" : 45.4905548095703,
         "id" : 24758350,
         "address_1" : "via Venini 42",
         "repinned" : false,
         "lon" : 9.21537780761719
      },
      "time" : 1481824800000,
      "id" : "235776931",
      "utc_offset" : 3600000,
      "waitlist_count" : 0,
      "created" : 1479902994000,
      "link" : "https://www.meetup.com/it-IT/Milano-Chatbots-Meetup/events/235776931/",
      "updated" : 1479903171000,
      "group" : {
         "created" : 1472632036000,
         "lon" : 9.1899995803833,
         "name" : "Milano Chatbots Meetup",
         "who" : "Chatbotters",
         "id" : 20372992,
         "lat" : 45.4599990844727,
         "urlname" : "Milano-Chatbots-Meetup",
         "join_mode" : "open"
      },
      "visibility" : "public",
      "yes_rsvp_count" : 68
   },
   {
      "name" : "Altro titolo",
      "how_to_find_us" : "MM1 Pasteur o tram 1 fermata Venini/Sauli. Citofono Mikamai/LinkMe/Venini42.",
      "status" : "upcoming",
      "description" : "una descrizione",
      "venue" : {
         "name" : "Mikamai e LinkMe",
         "city" : "Milano",
         "country" : "it",
         "localized_country_name" : "Italia",
         "lat" : 45.4905548095703,
         "id" : 24758350,
         "address_1" : "via Venini 42",
         "repinned" : false,
         "lon" : 9.21537780761719
      },
      "time" : 1481824800000,
      "id" : "235776931",
      "utc_offset" : 3600000,
      "waitlist_count" : 0,
      "created" : 1479902994000,
      "link" : "https://www.meetup.com/it-IT/Milano-Chatbots-Meetup/events/235776931/",
      "updated" : 1479903171000,
      "group" : {
         "created" : 1472632036000,
         "lon" : 9.1899995803833,
         "name" : "Milano Chatbots Meetup",
         "who" : "Chatbotters",
         "id" : 20372992,
         "lat" : 45.4599990844727,
         "urlname" : "Milano-Chatbots-Meetup",
         "join_mode" : "open"
      },
      "visibility" : "public",
      "yes_rsvp_count" : 68
   }
]
`
