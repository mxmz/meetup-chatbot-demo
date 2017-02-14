package thebot

import (
	"log"
	. "telegram-meetup/types"

	"golang.org/x/net/context"
)

func init() {
	registerCommandHandle("list", (*theBot).handleListCommand)
}

func (b *theBot) handleListCommand(ctx context.Context, m *InboundChatMessage) error {
	props, err := b.users.GetUserProperties(ctx, UserID(m.SenderID), "MeetupTokens")
	if err != nil {
		return err
	}
	if props == nil || len(props["MeetupTokens"]) == 0 {
		return b.startAuthorizationProcess(ctx, m)
	} else {
		tokens := MeetupAuthTokens(props["MeetupTokens"])
		groups, err := b.meetups.GetGroupList(ctx, tokens)
		if err != nil {
			if err == ErrMeetup401 {
				return b.startAuthorizationProcess(ctx, m)
			} else {
				return err
			}

		}
		buttons := make([]Button, 0, len(groups))
		for _, g := range groups {
			buttons = append(buttons, Button{
				Text:    g.Name,
				Command: "next " + g.Urlname,
			})
		}
		b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: m.SenderID,
			Message:     "️️️Ecco i tuoi gruppi:",
			Buttons:     buttons,
		})
	}
	return nil
}

func (b *theBot) HandleMeetupAuthCompletion(ctx context.Context, mc *MeetupAuthCompletion) error {
	var m InboundChatMessage
	m.UnmarshalMsg(mc.State)
	log.Println(m)
	b.users.SetUserProperties(ctx, UserID(m.SenderID), KVMap{
		"MeetupTokens": Value(mc.Tokens),
	})
	b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
		RecipientID: m.SenderID,
		Message:     "️️️Ok. Grazie per aver concesso l'autorizzazione. 😊",
	})
	return b.HandleInboundChatMessage(ctx, &m)
}
