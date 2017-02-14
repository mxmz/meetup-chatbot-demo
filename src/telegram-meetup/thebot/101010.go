package thebot

import (
	"fmt"
	"log"
	"strings"
	. "telegram-meetup/types"

	"golang.org/x/net/context"
)

type bot101010 struct {
	env     Env
	users   Repository
	queue   MessageEnqueuer
	meetups MeetupService
}

func (b *bot101010) HandleInboundChatMessage(ctx context.Context, m *InboundChatMessage) error {
	if len(m.Command) > 0 {
		if m.Command[0] == "start" {
			return b.handleStartCommand(ctx, m, "Buonasera, "+m.SenderName+".")
		}
		if m.SenderID == b.env.Get("ADMIN_CHAT_ID") {
			if m.Command[0] == "b" && len(m.Command) > 1 {
				return b.broadcastMessage(ctx, fmt.Sprintf("%v", strings.Join(m.Command[1:], " ")))
			}
		}
	}
	if len(m.Message) > 0 {
		if strings.Contains(m.Message, "?") {
			return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: b.env.Get("ADMIN_CHAT_ID"),
				Message:     "❓ [" + m.SenderID + "] " + m.SenderName + ":\n«" + m.Message + "»",
			})
		}
	}
	return nil
}

func (b *bot101010) HandleMeetupAuthCompletion(ctx context.Context, mc *MeetupAuthCompletion) error {
	return nil
}

func (b *bot101010) HandleJobRequest(ctx context.Context, m *JobRequest) error {
	log.Println("JobRequest: ----------> ", m)
	if len(m.Command) > 1 && m.Command[0] == "broadcast" {
		return b.broadcastMessage(ctx, fmt.Sprintf("%v", strings.Join(m.Command[1:], " ")))
	}
	return nil
}

func (b *bot101010) broadcastMessage(ctx context.Context, s string) error {

	outC, errC := b.users.FindMeetupUserLinksNotMatch(ctx, MeetupID(MyMeetupGroup),
		"Tag", Value("_"),
	)

	log.Println("reading results ...")

	for v := range outC {
		props, _ := b.users.GetUserProperties(ctx, v.UserID, "ID", "Name")
		log.Println("<<<<<<<", v, props)
		b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: string(v.UserID),
			Message:     s,
		},
		)
	}
	err := <-errC
	log.Println(err)

	return nil
}

func (b *bot101010) handleStartCommand(ctx context.Context, m *InboundChatMessage, response string) error {

	var userid = UserID(m.SenderID)
	props, err := b.users.GetUserProperties(ctx, userid, "ID", "Name")
	if err != nil {
		return err
	} else {
		if props == nil {
			props := KVMap{
				"ID":   Value(m.SenderID),
				"Name": Value(m.SenderName),
			}
			b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: b.env.Get("ADMIN_CHAT_ID"),
				Message:     "🆕 👤 [" + m.SenderID + "] " + m.SenderName,
			})
			err = b.users.SetUserProperties(ctx, userid, props)
			if err != nil {
				return err
			}
		} else {

		}
		err = b.users.SetMeetupUserLinkProperties(ctx, MeetupUserLink{MeetupID(MyMeetupGroup), userid}, KVMap{
			"Tag": Value(EmptyTag)},
		)
	}
	if err == nil {
		err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: string(userid),
			Message:     response,
		})

	}
	return nil
}
