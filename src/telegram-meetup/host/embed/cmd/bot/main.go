package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-meetup/app"
	"telegram-meetup/telegram"
	"telegram-meetup/tskbroker"
	"telegram-meetup/types"
	"time"

	"golang.org/x/net/context"

	"telegram-meetup/env"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pborman/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

const TelegramWebhook = "/telegram-webhook"
const FacebookWebhook = "/facebook-webhook"

func main() {
	var mq = tskbroker.NewMqBrokerMap(context.TODO(), 10, 10)
	go func() {
		appHost := NewAppHost(mq)
		app.RegisterInboundWorker(appHost)
		appHost.Run(context.TODO())
	}()
	go func() {
		appHost := NewAppHost(mq)
		app.RegisterOutboundWorker(appHost)
		appHost.Run(context.TODO())
	}()
	go func() {
		appHost := NewAppHost(mq)
		app.RegisterChatBotApp(appHost)
		appHost.Run(context.TODO())
	}()

	/*
		go func() {
			for {
				mq.Listen(context.TODO(), func(m *tskbroker.MqMessage) error {

					msg := m.Payload.(*types.WebhookMessage)
					log.Println(msg.Type)
					log.Println(string(msg.Data))

					return nil
				}, WebhookMessageName)
			}
		}()
	*/

	startHttp(mq)
}

func startHttp(mq *tskbroker.MqBrokerMap) {
	env := env.LoadEnv()

	http.HandleFunc(TelegramWebhook, func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(461)
			fmt.Fprint(w, err)
		} else {
			appHost := NewAppHost(mq)
			appHost.handleWebhookMessage(r.Context(), &types.WebhookMessage{"telegram", body})
			fmt.Fprint(w, "ok")
		}
	})
	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		proto := nvl(r.Header.Get("X-Forwarded-Proto"), "https")
		hostname := nvl(r.Host, r.Header.Get("Host"))
		log.Println(proto)
		log.Println(hostname)
		url := proto + "://" + hostname
		bot, _ := telegram.NewBot(E, nil, func(context.Context) *http.Client {
			return http.DefaultClient
		})
		username, err := bot.SetWebhook(r.Context(), url+TelegramWebhook)

		botUrl := "http://telegram.me/" + username

		log.Println(err)
		fmt.Fprint(w, url, "\n")

		png, err := qrcode.Encode(botUrl, qrcode.Medium, 800)
		msg := &tskbroker.MqMessage{
			ID:      uuid.New(),
			Type:    OutboundMessageQueueName,
			Delayed: time.Now(),
			Payload: &types.OutboundChatMessage{RecipientID: env.Get("ADMIN_CHAT_ID"),
				Message:    "Eccomi qui. " + botUrl,
				Attachment: &types.File{username + ".png", png},
			},
		}
		log.Println("inject", msg)
		mq.Inject(r.Context(), OutboundMessageQueueName, msg)

	})

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(200)
			ctx := r.Context()

			bot, _ := telegram.NewBot(env, nil, func(ctx context.Context) *http.Client {
				return http.DefaultClient
			})
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
	http.HandleFunc("/qrcode.png",
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(200)
			ctx := r.Context()

			bot, _ := telegram.NewBot(env, nil, func(ctx context.Context) *http.Client {
				return http.DefaultClient
			})
			api, _ := bot.MakeApi(ctx)
			me, _ := api.GetMe()

			botUrl := "http://telegram.me/" + me.UserName + "?start=qrcode"

			png, _ := qrcode.Encode(botUrl, qrcode.Medium, 500)
			w.Write(png)
		},
	)

	log.Println("starting http server ....")
	srv := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         "0.0.0.0:8080",
	}
	log.Fatal(srv.ListenAndServe())
}
