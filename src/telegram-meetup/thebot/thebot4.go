package thebot

import (
	. "telegram-meetup/types"
	"time"

	"golang.org/x/net/context"
)

func init() {
	registerCommandHandle("vote_default", (*theBot).handleVoteDefaultCommand)
	registerCommandHandle("poll_default", func(b *theBot, ctx context.Context, m *InboundChatMessage) error {
		pollButtons := []Button{}
		for k, v := range b.defaultPoll {
			pollButtons = append(pollButtons,
				Button{v, "vote_default " + k},
			)
		}
		return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: m.SenderID,
			Message:     "Vota il migliore Default:",
			Buttons:     pollButtons,
		})

	})
}

func (b *theBot) handleVoteDefaultCommand(ctx context.Context, m *InboundChatMessage) error {
	if len(m.Command) < 2 || len(m.Command[1]) < 1 {
		return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: m.SenderID,
			Message:     "️️️Errore: manca argomento",
		})
	}
	var vote = m.Command[1]

	b.users.SetPollUserLinkProperties(ctx,
		PollUserLink{PollID("Default"), UserID(m.SenderID)},
		KVMap{"Vote": Value(vote), "Tag": Value(EmptyTag)},
	)
	b.queue.HandleJobRequest(ctx, &JobRequest{
		Cookie:    []byte("cookie"),
		Command:   []string{"notify_poll_results", "Default"},
		Scheduled: time.Now().Add(2 * time.Second),
		Droppable: time.Now().Add(120 * time.Second),
	})

	return nil
}
