package thebot

import (
	"log"
	"strings"
	. "telegram-meetup/types"
	"time"

	"golang.org/x/net/context"
)

func init() {
	registerCommandHandle("attending", (*theBot).handleAttending)
	registerCommandHandle("not_attending", (*theBot).handleNotAttending)
	registerCommandHandle("bce", (*theBot).handleBroadcastCurrentEvent)
}

func (b *theBot) handleAttending(ctx context.Context, m *InboundChatMessage) error {
	log.Println(m)
	if len(m.Command) > 1 { // /attending <etag>
		group := MyMeetupGroup
		userid := m.SenderID
		next, err := b.meetups.GetNextMeetup(ctx, MeetupGroupID(group), time.Now())
		log.Println(err)
		if err == nil && next != nil && eventInProgress(next) {
			err = b.users.SetMeetupUserLinkProperties(
				ctx,
				MeetupUserLink{MeetupID(group), UserID(userid)}, KVMap{
					"PTag":  Value(next.ETag),
					"PTime": Value(time.Now().String()),
				})
			b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: m.SenderID,
				Message:     "Ok. Buona serata. 😊",
				Buttons:     []Button{},
			})
			if len(b.adminChatID) > 0 {
				b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
					RecipientID: b.adminChatID,
					Message:     "👤 " + m.SenderID + "\n" + m.SenderName + " è presente.",
					Buttons:     []Button{},
				})
			}
		}
		return err
	}
	return nil
}

func (b *theBot) handleNotAttending(ctx context.Context, m *InboundChatMessage) error {
	log.Println(m)
	if len(m.Command) > 1 { //  /not_attending <etag>
		group := MyMeetupGroup
		userid := m.SenderID
		err := b.users.SetMeetupUserLinkProperties(
			ctx,
			MeetupUserLink{MeetupID(group), UserID(userid)}, KVMap{
				"PTag":  Value(EmptyTag),
				"PTime": Value(time.Now().String()),
			})
		b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: m.SenderID,
			Message:     "OK.",
			Buttons:     []Button{},
		})
		return err
	}
	return nil
}

func (b *theBot) handleBroadcastCurrentEvent(ctx context.Context, m *InboundChatMessage) error {
	if len(b.adminChatID) > 0 && m.SenderID == b.adminChatID && len(m.Command) > 1 {
		return b.queue.HandleJobRequest(ctx, &JobRequest{
			Cookie:    []byte(""),
			Command:   []string{"broadcast_current_event", strings.Join(m.Command[1:], " ")},
			Scheduled: time.Now(),
			Droppable: time.Now().Add(120 * time.Minute),
		})
	}
	return nil
}

func (b *theBot) broadcastMessageCurrentEvent(ctx context.Context, s string) error {
	next, err := b.meetups.GetNextMeetup(ctx, MeetupGroupID(MyMeetupGroup), time.Now())
	log.Println(err)
	if err == nil && next != nil {
		outC, errC := b.users.FindMeetupUserLinksMatch(ctx, MeetupID(MyMeetupGroup),
			"PTag", Value(next.ETag),
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
	}
	return nil
}
