package thebot

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	. "telegram-meetup/types"
	"time"

	"golang.org/x/net/context"

	"encoding/json"
	"strings"

	"github.com/glycerine/zebrapack/msgp"
	"github.com/soh335/ical"
)

const MyMeetupGroup = "Milano-Chatbots-Meetup"

type theBot struct {
	users       Repository
	queue       MessageEnqueuer
	meetups     MeetupService
	meetupAuth  MeetupAuthorizer
	adminChatID string
	defaultPoll map[string]string
}

const EmptyTag = "EMPTY"

func NewBot(env Env, db Repository, q MessageEnqueuer, m MeetupService, ma MeetupAuthorizer) App {
	var defaultPollJson = env.Get("DEFAULT_POLL")
	var defaultPoll map[string]string
	if len(defaultPollJson) > 8 {
		json.Unmarshal([]byte(defaultPollJson), &defaultPoll)
		log.Println(" ----------> ", defaultPoll)
	}
	if env.Get("BOT_MODE") == "101010" {
		return &bot101010{env, db, q, m}
	} else {
		return &theBot{db, q, m, ma, env.Get("ADMIN_CHAT_ID"), defaultPoll}
	}

}

func (b *theBot) handleStartCommand(ctx context.Context, m *InboundChatMessage) error {
	var welcomeMsg = "Benvenuto, " + m.SenderName + "."
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
			err = b.users.SetUserProperties(ctx, userid, props)
			if err != nil {
				return err
			}
		} else {

		}
		err = b.users.SetMeetupUserLinkProperties(ctx, MeetupUserLink{MeetupID(MyMeetupGroup), userid}, KVMap{
			"Tag": Value(EmptyTag),
		})
		if err == nil {
			err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: string(userid),
				Message:     welcomeMsg,
				Buttons: []Button{
					Button{"ℹ️ Prossimo", "next"},
					Button{"🌐 Mappa", "map"},
				},
			})
			b.queue.HandleJobRequest(ctx, &JobRequest{
				Cookie:    []byte("cookie"),
				Command:   []string{"do", "this"},
				Scheduled: time.Now().Add(20 * time.Second),
				Droppable: time.Now().Add(120 * time.Second),
			})
			var group = MyMeetupGroup
			next, err := b.meetups.GetNextMeetup(ctx, MeetupGroupID(group), time.Now())
			if err == nil && next != nil && eventInProgress(next) {
				err = b.notifyMeetup(ctx, m.SenderID, group, next)
				if err == nil {
					err = b.users.SetMeetupUserLinkProperties(ctx, MeetupUserLink{MeetupID(group), userid}, KVMap{
						"Tag": Value(next.ETag),
					})
				}
			}

			if len(b.adminChatID) > 0 {
				b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
					RecipientID: b.adminChatID,
					Message:     "👤 " + m.SenderID + "\n" + m.SenderName,
					Buttons:     []Button{},
				})
			}

		}
		return err

	}

}
func (b *theBot) handleNextCommand(ctx context.Context, m *InboundChatMessage) error {
	var userid = UserID(m.SenderID)
	props, err := b.users.GetUserProperties(ctx, userid, "Name")
	if err != nil {
		return err
	} else {
		if props == nil {
			return errors.New("user not found: " + m.SenderID)
		}
		var group = MyMeetupGroup
		if len(m.Command) > 1 {
			group = m.Command[1]
		}
		next, err := b.meetups.GetNextMeetup(ctx, MeetupGroupID(group), time.Now())
		if err == nil {
			if next != nil {
				err = b.notifyMeetup(ctx, m.SenderID, group, next)
				if err == nil {
					err = b.users.SetMeetupUserLinkProperties(ctx, MeetupUserLink{MeetupID(group), userid}, KVMap{
						"Tag": Value(next.ETag),
					})
				}
			} else {
				err = b.notifyNoMeetupScheduledYet(ctx, m.SenderID, group)
				if err == nil {
					err = b.users.SetMeetupUserLinkProperties(
						ctx,
						MeetupUserLink{MeetupID(group), userid}, KVMap{
							"Tag": Value(EmptyTag),
						})
				}

			}
		}
	}
	return err
}

func (b *theBot) handleMapCommand(ctx context.Context, m *InboundChatMessage) error {
	var userid = UserID(m.SenderID)
	props, err := b.users.GetUserProperties(ctx, userid, "Name")
	if err != nil {
		return err
	} else {
		if props == nil {
			return errors.New("user not found")
		}
		next, err := b.meetups.GetNextMeetup(ctx, MyMeetupGroup, time.Now())
		if err == nil {
			if next != nil {
				err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
					RecipientID: m.SenderID,
					Message:     fmt.Sprintf("📍 %s", next.Where),
					Map:         next.Map,
				})
			} else {
				err = b.notifyNoMeetupScheduledYet(ctx, m.SenderID, MyMeetupGroup)
			}
		}
	}
	return err

}

var commandHandlerMap = map[string]func(b *theBot, ctx context.Context, m *InboundChatMessage) error{}

func registerCommandHandle(name string, method func(*theBot, context.Context, *InboundChatMessage) error) {
	commandHandlerMap[name] = method
}

func init() {
	registerCommandHandle("start", (*theBot).handleStartCommand)
	registerCommandHandle("info", (*theBot).handleNextCommand)
	registerCommandHandle("next", (*theBot).handleNextCommand)
	registerCommandHandle("map", (*theBot).handleMapCommand)
}

func (b *theBot) HandleInboundChatMessage(ctx context.Context, m *InboundChatMessage) error {
	if len(m.Command) > 0 {
		var impl, ok = commandHandlerMap[m.Command[0]]
		if ok {
			return impl(b, ctx, m)
		} else if m.Command[0] == "start+next" {
			err := b.handleStartCommand(ctx, m)
			if err == nil {
				m.Command[0] = "next"
				return b.HandleInboundChatMessage(ctx, m)
			} else {
				return err
			}
		} else if m.Command[0] == "cal" {
			var userid = UserID(m.SenderID)
			props, err := b.users.GetUserProperties(ctx, userid, "Name")
			if err != nil {
				return err
			} else {
				if props == nil {
					return errors.New("user not found")
				}
				next, err := b.meetups.GetNextMeetup(ctx, MyMeetupGroup, time.Now())
				if err == nil {
					if next != nil {

						err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
							RecipientID: m.SenderID,
							Message:     fmt.Sprintf("💬 %s", next.Title),
							Attachment:  IcalFromEvent(MyMeetupGroup, next),
						})

					} else {
						err = b.notifyNoMeetupScheduledYet(ctx, m.SenderID, MyMeetupGroup)
					}
				}
			}
			return err
		} else if m.Command[0] == "vote_default" {
			return b.handleVoteDefaultCommand(ctx, m)
		} else if m.Command[0] == "poll_default" {
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
		} else {
			//return errors.New("unexpected message")
			log.Println("ignored: " + m.Command[0])
		}
	}
	log.Println(m.Message)
	if len(m.Message) > 0 && len(b.adminChatID) > 0 {
		return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: b.adminChatID,
			Message:     "💬 [" + m.SenderID + "] " + m.SenderName + ":\n«" + m.Message + "»",
		})
	}
	return nil
}

func (b *theBot) startAuthorizationProcess(ctx context.Context, m *InboundChatMessage) error {
	buf := bytes.Buffer{}
	wrt := msgp.NewWriter(&buf)
	err := m.EncodeMsg(wrt)
	wrt.Flush()
	u, err := b.meetupAuth.MakeOOBAuthStartURL(ctx, buf.Bytes())
	if err != nil {
		log.Println(err)
		return err
	}
	err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
		RecipientID: m.SenderID,
		Message:     "Segui il link per consentire al bot di accedere a Meetup:\n" + string(u),
	})
	return err
}

func (b *theBot) notifyNoMeetupScheduledYet(ctx context.Context, userid string, meetup string) error {
	return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
		RecipientID: userid,
		Message:     "️️➡️ " + meetup + "\nNessun incontro in programma. 😞",
	},
	)

}

func (b *theBot) notifyMeetup(ctx context.Context, userid string, meetup string, event *Meetup) error {
	err := b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
		RecipientID: userid,
		Message: fmt.Sprintf(
			"️➡️ %s\n"+
				"💬 %s\n"+
				"📍 %s\n"+
				"📆 %v\n",
			meetup,
			event.Title,
			event.Where,
			event.When,
		),
	})

	if err == nil && eventInProgress(event) {
		err = b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
			RecipientID: string(userid),
			Message:     "💬 L'evento è in corso. Stai partecipando?",
			Buttons: []Button{
				Button{"✔️ Sì", "attending " + event.ETag},
				Button{"❌ No", "not_attending " + event.ETag},
			},
		})
	}
	return err
}

func IcalFromEvent(meetup string, event *Meetup) *File {
	vComponents := []ical.VComponent{
		&ical.VEvent{
			UID:         "123",
			DTSTAMP:     event.When,
			DTSTART:     event.When,
			DTEND:       event.When.Add(2 * time.Hour),
			SUMMARY:     event.Title,
			DESCRIPTION: event.Title,
			TZID:        "Europe/Rome",
		},
	}

	cal := ical.NewBasicVCalendar()
	cal.PRODID = "proid"
	cal.X_WR_TIMEZONE = "Europe/Rome"
	cal.X_WR_CALNAME = "Meetup"
	cal.X_WR_CALDESC = "Meeting"

	cal.VComponent = vComponents

	var buf bytes.Buffer

	cal.Encode(&buf)
	filename := strings.Map(func(c rune) rune {
		if (c >= 'A' && c <= 'Z') || c >= 'a' && c <= 'z' {
			return c
		} else {
			return '_'
		}

	}, meetup) + ".ics"

	return &File{filename, []byte(buf.Bytes())}

}

func (b *theBot) broadcastMessage(ctx context.Context, s string) error {

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

func (b *theBot) notifyNewMeetup(ctx context.Context, meetup string) error {
	tag := EmptyTag
	notify := func(userid string) error {
		return b.notifyNoMeetupScheduledYet(ctx, userid, meetup)
	}
	next, err := b.meetups.GetNextMeetup(ctx, MeetupGroupID(meetup), time.Now())
	if err != nil {
		return err
	} else {
		if next != nil {
			tag = next.ETag
			notify = func(userid string) error {
				return b.notifyMeetup(ctx, userid, meetup, next)
			}
		}
		outC, errC := b.users.FindMeetupUserLinksNotMatch(ctx, MeetupID(meetup), "Tag", Value(tag))
		log.Println("read outC ...")
		for v := range outC {
			log.Println("<<<<", v)
			err := notify(string(v.UserID))
			if err == nil {
				err = b.users.SetMeetupUserLinkProperties(
					ctx,
					MeetupUserLink{v.MeetupID, v.UserID}, KVMap{
						"Tag": Value(tag),
					})
			}
		}
		err := <-errC
		log.Println(err)
	}
	return nil
}

var emojiNumbers = map[int]string{
	0: "1️⃣",
	1: "️2️⃣",
	2: "3️⃣",
	3: "4️⃣",
	4: "5️⃣",
}

func (b *theBot) notifyNewPollResults(ctx context.Context, poll string) error {
	tag := EmptyTag
	notify := func(userid string) error {
		return nil
	}

	res, err := b.users.GetPollResult(ctx, PollID(poll), Key("Vote"))
	if err != nil {
		return err
	} else {

		pollButtons := []Button{}

		for k, v := range b.defaultPoll {
			pollButtons = append(pollButtons,
				Button{fmt.Sprintf("%s: %d\n", v, res[k]), "vote_default " + k},
			)
		}

		tag = md5hex(fmt.Sprintf("%v", res))

		notify = func(userid string) error {
			return b.queue.HandleOutboundChatMessage(ctx, &OutboundChatMessage{
				RecipientID: userid,
				Message:     "Risultati: ",
				Buttons:     pollButtons,
			})
		}
		outC, errC := b.users.FindPollUserLinksNotMatch(ctx, PollID(poll), "Tag", Value(tag))
		log.Println("read outC ...")
		for v := range outC {
			log.Println("<<<<", v)
			err := notify(string(v.UserID))
			if err == nil {
				err = b.users.SetPollUserLinkProperties(
					ctx,
					PollUserLink{v.PollID, v.UserID}, KVMap{
						"Tag": Value(tag),
					})
			}
		}
		err = <-errC
		log.Println(err)
	}
	return nil
}

func (b *theBot) HandleJobRequest(ctx context.Context, m *JobRequest) error {
	log.Println("JobRequest: ----------> ", m)
	if len(m.Command) > 1 && m.Command[0] == "notify_poll_results" {
		return b.notifyNewPollResults(ctx, "Default")
	} else if len(m.Command) > 1 && m.Command[0] == "broadcast" {
		return b.broadcastMessage(ctx, fmt.Sprintf("%v", strings.Join(m.Command[1:], " ")))

	} else if len(m.Command) > 1 && m.Command[0] == "broadcast_current_event" {
		return b.broadcastMessageCurrentEvent(ctx, fmt.Sprintf("%v", strings.Join(m.Command[1:], " ")))
	}
	if len(m.Command) > 0 && m.Command[0] == "check_new_notify" {
		return b.notifyNewMeetup(ctx, MyMeetupGroup)
	}
	return nil
}
