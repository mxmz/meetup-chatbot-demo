package telegram

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	. "telegram-meetup/types"

	tgbot "gopkg.in/telegram-bot-api.v4"

	"errors"
	"log"
	"strings"
)

type clientFunc func(context.Context) *http.Client

type bot struct {
	client  clientFunc
	token   string
	inqueue InboundChatMessageHandler
}

/*

 */

func NewBot(env Env, inqueue InboundChatMessageHandler, c clientFunc) (*bot, error) {
	var token = env.Get("TELEGRAM_TOKEN")
	if len(token) < 40 {
		return nil, errors.New("undefined or invalid Telegram token: " + token)
	}
	return &bot{c, token, inqueue}, nil
}

func (b *bot) MakeApi(ctx context.Context) (*tgbot.BotAPI, error) {
	api, err := tgbot.NewBotAPIWithClient(b.token, b.client(ctx))
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (b *bot) HandleWebhookMessage(ctx context.Context, m *WebhookMessage) error {
	//log.Println(m, b)
	if m.Type != "telegram" {
		return errors.New("invalid message type")
	}
	return b.HandleWebhookPayload(ctx, m.Data)
}

func (b *bot) HandleWebhookPayload(ctx context.Context, bytes []byte) error {
	log.Println(string(bytes))
	update := tgbot.Update{}
	err := json.Unmarshal(bytes, &update)
	if err != nil {
		return err
	}

	var msg InboundChatMessage

	var from *tgbot.User
	var message string
	var command string

	if update.CallbackQuery != nil {
		api, err := b.MakeApi(ctx)
		if err != nil {
			return err
		}
		from = update.CallbackQuery.From
		message = ""
		command = update.CallbackQuery.Data
		api.AnswerCallbackQuery(tgbot.CallbackConfig{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            "⚙ /" + command + " ... ",
			ShowAlert:       false,
		})

	} else {

		from = update.Message.From
		message = update.Message.Text
		command = getCommandFromUpdate(&update)

		if strings.HasPrefix(message, "/"+command+" ") {
			command = message[1:]
		}

	}

	msg.SenderID = "telegram:" + strconv.FormatInt(int64(from.ID), 10)
	msg.SenderName = strings.Join([]string{from.FirstName, from.LastName}, " ")
	if len(from.UserName) > 0 {
		msg.SenderName += " (@" + from.UserName + ")"
	}
	if len(command) > 0 {
		msg.Command = strings.Split(command, " ")
	}
	msg.Message = message
	msgs := []*InboundChatMessage{&msg}
	if err == nil {
		for _, m := range msgs {
			log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", m)
			err := b.inqueue.HandleInboundChatMessage(ctx, m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *bot) HandleOutboundChatMessage(ctx context.Context, m *OutboundChatMessage) error {
	return b.PostMessage(ctx, m)
}

func (b *bot) PostMessage(ctx context.Context, m *OutboundChatMessage) error {
	var recipient = m.RecipientID
	var parsedRec = strings.SplitN(recipient, ":", 2)
	if len(parsedRec) < 2 || parsedRec[0] != "telegram" {
		return errors.New("invalid recipient id: " + recipient)
	}
	recipient = parsedRec[1]
	text := m.Message
	chatID, err := strconv.ParseInt(recipient, 10, 64)
	if err != nil {
		return err
	} else {
		api, err := b.MakeApi(ctx)
		if err != nil {
			return err
		}
		tm := tgbot.NewMessage(int64(chatID), text)
		tm.DisableWebPagePreview = true
		if m.Buttons != nil && len(m.Buttons) > 0 {
			log.Println(m.Buttons)
			bs := [][]tgbot.InlineKeyboardButton{}
			for _, b := range m.Buttons {
				bs = append(bs, []tgbot.InlineKeyboardButton{tgbot.NewInlineKeyboardButtonData(b.Text, b.Command)})
			}
			kb := tgbot.NewInlineKeyboardMarkup(bs...)
			tm.ReplyMarkup = kb
		}
		_, err = api.Send(tm)
		if err == nil && m.Map != nil {
			mm := tgbot.NewLocation(int64(chatID), m.Map.Lat, m.Map.Lng)
			_, err = api.Send(mm)
		}
		if m.Attachment != nil && len(m.Attachment.Data) > 0 {
			data := tgbot.FileBytes{m.Attachment.Name, m.Attachment.Data}
			file := tgbot.NewDocumentUpload(int64(chatID), data)
			api.Send(file)
		}
		//tm2 := tgbot.NewMessage(int64(chatID), "test")
		//kb := tgbot.NewInlineKeyboardMarkup([]tgbot.InlineKeyboardButton{
		//	tgbot.NewInlineKeyboardButtonData("ℹ️ Informazioni", "info"),
		//	tgbot.NewInlineKeyboardButtonData("🌐 Indicazioni", "map"),
		//})
		//kb := tgbot.NewReplyKeyboard([]tgbot.KeyboardButton{
		//	tgbot.NewKeyboardButton("/map"),
		//})
		//tm2.ReplyMarkup = kb
		//_, err = b.api.Send(tm2)
		return err
	}
}

func getCommandFromUpdate(update *tgbot.Update) string {
	if update.Message != nil && update.Message.Entities != nil && (*update.Message.Entities)[0].Type == "bot_command" {
		e := (*update.Message.Entities)[0]
		return update.Message.Text[e.Offset+1 : e.Offset+e.Length]
	} else {
		return "nop"
	}
}

func (b *bot) SetWebhook(ctx context.Context, url string) (string, error) {
	api, err := b.MakeApi(ctx)
	if err != nil {
		return "", err
	}
	c := tgbot.NewWebhook(url)
	log.Println(c)
	_, err = api.SetWebhook(c)
	log.Println(api.GetMe())
	user, err := api.GetMe()
	return user.UserName, err
}

func nvl(s1, s2 string) string {
	if len(s1) == 0 {
		return s2
	}
	return s1
}
