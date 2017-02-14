package app

import (
	"net/http"
	"telegram-meetup/telegram"
	"telegram-meetup/thebot"
	. "telegram-meetup/types"

	"golang.org/x/net/context"
)

func RegisterInboundWorker(host AppHost) {
	makeClient := func(ctx context.Context) *http.Client {
		return host.MakeHttpClient(ctx)
	}
	tg, _ := telegram.NewBot(host.MakeEnv(), host, makeClient)
	host.RegisterWebhookMessageHandler(tg)
}

func RegisterChatBotApp(host AppHost) {
	bot := thebot.NewBot(host.MakeEnv(), host.MakeRepository(), host, host.MakeMeetupService(), host.MakeMeetupAuthorizer())
	host.RegisterInboundChatMessagehandler(bot)
	host.RegisterJobRequesteHandler(bot)
}

func RegisterOutboundWorker(host AppHost) {
	makeClient := func(ctx context.Context) *http.Client {
		return host.MakeHttpClient(ctx)
	}
	tg, _ := telegram.NewBot(host.MakeEnv(), host, makeClient)
	host.RegisterOutboundChatMessagehandler(tg)

}
