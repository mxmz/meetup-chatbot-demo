package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"telegram-meetup/app"
	"telegram-meetup/telegram"
	"telegram-meetup/types"
	"time"

	"golang.org/x/net/context"

	"telegram-meetup/env"

	_ "github.com/mattn/go-sqlite3"
	qrcode "github.com/skip2/go-qrcode"
)

const TelegramWebhook = "/telegram-webhook"

func main() {

	go func() {
		appHost := NewAppHost()
		app.RegisterChatBotApp(appHost)
		app.RegisterInboundWorker(appHost)
		app.RegisterOutboundWorker(appHost)
		appHost.Run(context.TODO())
	}()

	startHttp()
}

func startHttp() {
	env := env.LoadEnv()

	http.HandleFunc(TelegramWebhook, func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(461)
			fmt.Fprint(w, err)
		} else {
			appHost := NewAppHost()
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

		msg := &types.OutboundChatMessage{RecipientID: env.Get("ADMIN_CHAT_ID"),
			Message:    "Eccomi qui. " + botUrl,
			Attachment: &types.File{username + ".png", png},
		}
		log.Println("inject", msg)
		appHost := NewAppHost()
		err = appHost.injectMessage(r.Context(), OutboundMessageQueueName, msg, time.Now())
		fmt.Fprint(w, err, "\n")
	})

	log.Println("starting http server ....")
	srv := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         "127.0.0.1:8080",
	}
	log.Fatal(srv.ListenAndServe())
}
