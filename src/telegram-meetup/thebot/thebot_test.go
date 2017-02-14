package thebot

import (
	"log"
	"sort"
	"telegram-meetup/types"
	. "telegram-meetup/types"
	"testing"
	"time"

	"golang.org/x/net/context"

	"strings"

	"os"

	. "github.com/onsi/gomega"
)

var meetupTime, _ = time.Parse("Jan 2 15:04:05 -0700 MST 2006", "Jan 3 15:04:05 -0700 MST 2106")

func TestNewBot(t *testing.T) {
	RegisterTestingT(t)

	var b = NewBot(&osEnv{}, &mockRepository{}, &mockMsgQueue{}, &mockMeetupService{}, &mockMeetupService{})

	Expect(b).To(BeEquivalentTo(&theBot{&mockRepository{}, &mockMsgQueue{}, &mockMeetupService{}, &mockMeetupService{}, "", nil}))
}

func Test_theBot_HandleInboundChatMessage(t *testing.T) {
	RegisterTestingT(t)
	var r = &mockRepository{}
	var mq = &mockMsgQueue{}
	var ms = &mockMeetupService{}
	var b = NewBot(&osEnv{}, r, mq, ms, ms)
	t.Run("start + new user", func(t *testing.T) {
		b.HandleInboundChatMessage(context.TODO(), &InboundChatMessage{
			SenderID:   "u42",
			SenderName: "Mr 42",
			Command:    []string{"start"},
			Message:    "",
		})

		Expect(r.meetupUsers).To(BeEquivalentTo(map[string]map[string]string{
			"Milano-Chatbots-Meetup+u42": map[string]string{
				"Tag": "EMPTY",
			},
		}))
		Expect(r.users).To(BeEquivalentTo(map[string]map[string]string{
			"u42": map[string]string{
				"Name": "Mr 42",
				"ID":   "u42",
			},
		}))

		Expect(r.l.trace).To(BeEquivalentTo([]interface{}{
			"GetUserProperties",
			UserID("u42"),
			[]Key{"ID", "Name"},
			"SetUserProperties",
			UserID("u42"),
			KVMap{"ID": "u42", "Name": "Mr 42"},
			"SetMeetupUserLinkProperties",
			MeetupUserLink{MeetupID("Milano-Chatbots-Meetup"), UserID("u42")},
			KVMap{Key("Tag"): Value("EMPTY")},
		}))

		Expect(mq.l.trace).To(BeEquivalentTo([]interface{}{
			"HandleOutboundChatMessage",
			context.TODO(),
			&OutboundChatMessage{
				RecipientID: "u42",
				Message:     "Benvenuto, Mr 42.",
				Map:         nil,
				Buttons: []Button{
					Button{
						Text:    "ℹ️ Prossimo",
						Command: "next",
					},
					Button{Text: "🌐 Mappa", Command: "map"},
				},
				Attachment: nil,
			},

			//
		}))
	})

	r.l.reset()
	mq.l.reset()
	t.Run("start ~ existing user", func(t *testing.T) {
		b.HandleInboundChatMessage(context.TODO(), &InboundChatMessage{
			SenderID:   "u42",
			SenderName: "Mr 42",
			Command:    []string{"start"},
			Message:    "",
		})
		Expect(r.meetupUsers).To(BeEquivalentTo(map[string]map[string]string{
			"Milano-Chatbots-Meetup+u42": map[string]string{
				"Tag": "EMPTY",
			},
		}))
		Expect(r.users).To(BeEquivalentTo(map[string]map[string]string{
			"u42": map[string]string{
				"Name": "Mr 42",
				"ID":   "u42",
			},
		}))
		Expect(r.l.trace).To(BeEquivalentTo([]interface{}{
			"GetUserProperties",
			UserID("u42"),
			[]Key{"ID", "Name"},
			"SetMeetupUserLinkProperties",
			MeetupUserLink{MeetupID("Milano-Chatbots-Meetup"), UserID("u42")},
			KVMap{Key("Tag"): Value("EMPTY")},
		}))
	})
	r.l.reset()
	t.Run("start ~ new user", func(t *testing.T) {
		b.HandleInboundChatMessage(context.TODO(), &InboundChatMessage{
			SenderID:   "u43",
			SenderName: "Mr 43",
			Command:    []string{"start"},
			Message:    "",
		})

		Expect(r.meetupUsers).To(BeEquivalentTo(map[string]map[string]string{
			"Milano-Chatbots-Meetup+u42": map[string]string{
				"Tag": "EMPTY",
			},
			"Milano-Chatbots-Meetup+u43": map[string]string{
				"Tag": "EMPTY",
			},
		}))
		Expect(r.users).To(BeEquivalentTo(map[string]map[string]string{
			"u42": map[string]string{
				"Name": "Mr 42",
				"ID":   "u42",
			},
			"u43": map[string]string{
				"Name": "Mr 43",
				"ID":   "u43",
			},
		}))

		Expect(r.l.trace).To(BeEquivalentTo([]interface{}{
			"GetUserProperties",
			UserID("u43"),
			[]Key{"ID", "Name"},
			"SetUserProperties",
			UserID("u43"),
			KVMap{"ID": "u43", "Name": "Mr 43"},
			"SetMeetupUserLinkProperties",
			MeetupUserLink{MeetupID("Milano-Chatbots-Meetup"), UserID("u43")},
			KVMap{Key("Tag"): Value("EMPTY")},
		}))
	})

}

func Test_theBot_HandleScheduledTask(t *testing.T) {
}

type mockMsgQueue struct {
	enqueued     map[int]*OutboundChatMessage
	enqueueError error
	Calls        int
	l            logger
}

func (q *mockMsgQueue) HandleOutboundChatMessage(ctx context.Context, m *OutboundChatMessage) error {
	q.Calls++
	if q.enqueued == nil {
		q.enqueued = map[int]*OutboundChatMessage{}
	}
	q.enqueued[q.Calls] = m
	log.Println(m)
	q.l.write("HandleOutboundChatMessage", ctx, m)
	return q.enqueueError
}

func (q *mockMsgQueue) HandleJobRequest(ctx context.Context, m *JobRequest) error {
	return nil
}
func (q *mockMsgQueue) HandleInboundChatMessage(ctx context.Context, m *InboundChatMessage) error {
	return nil
}

type mockMeetupService struct {
	Exists bool
	When   time.Time
	GetErr error
	Calls  int
}

type osEnv struct {
}

func (*osEnv) Get(k string) string {
	return os.Getenv(k)
}

func (m *mockMeetupService) GetNextMeetup(ctx context.Context, name MeetupGroupID, ref time.Time) (*Meetup, error) {
	m.Calls++
	log.Println(m, "---->", m.Calls, m.Exists)
	if m.Exists {
		return &Meetup{
			ETag:  "4242424242424",
			When:  m.When,
			Where: "Milan",
			Title: "My nice bot",
		}, m.GetErr
	} else {
		return nil, m.GetErr
	}
}

func (m *mockMeetupService) MakeOOBAuthStartURL(context.Context, MeetupAuthState) (URL, error) {
	panic("not implemented")
}
func (m *mockMeetupService) RefreshAuthTokens(context.Context, MeetupAuthTokens) (MeetupAuthTokens, error) {
	panic("not implemented")
}
func (m *mockMeetupService) GetGroupList(context.Context, MeetupAuthTokens) ([]MeetupGroup, error) {
	panic("not implemented")
}

type logger struct {
	trace []interface{}
}

func (l *logger) reset() {
	l.trace = make([]interface{}, 0)
}

func (l *logger) write(as ...interface{}) {
	if l.trace == nil {
		l.reset()
	}
	for _, v := range as {
		l.trace = append(l.trace, v)
	}
}

type mockRepository struct {
	users       map[string]map[string]string
	meetupUsers map[string]map[string]string
	l           logger
}

func (r *mockRepository) SetUserProperties(ctx context.Context, id types.UserID, m types.KVMap) error {
	r.l.write("SetUserProperties", id, m)
	var u, ok = r.users[string(id)]
	if !ok {
		u = map[string]string{}
		r.users[string(id)] = u
	}
	for k, v := range m {
		u[string(k)] = string(v)
	}
	return nil
}

func (r *mockRepository) GetUserProperties(ctx context.Context, id types.UserID, ks ...types.Key) (types.KVMap, error) {
	r.l.write("GetUserProperties", id, ks)
	if r.users == nil {
		r.users = map[string]map[string]string{}
	}
	var u, ok = r.users[string(id)]
	if !ok {
		return nil, nil
	}
	var m = KVMap{}
	for _, k := range ks {
		var v, ok = u[string(k)]
		if ok {
			m[k] = Value(v)
		}
	}
	return m, nil
}

func (r *mockRepository) SetMeetupUserLinkProperties(ctx context.Context, l MeetupUserLink, m KVMap) error {
	r.l.write("SetMeetupUserLinkProperties", l, m)
	if r.meetupUsers == nil {
		r.meetupUsers = map[string]map[string]string{}
	}
	var id = string(l.MeetupID) + "+" + string(l.UserID)
	var mu, ok = r.meetupUsers[string(id)]
	if !ok {
		mu = map[string]string{}
		r.meetupUsers[string(id)] = mu
	}
	for k, v := range m {
		mu[string(k)] = string(v)
	}
	return nil

}

func (r *mockRepository) FindMeetupUserLinksNotMatch(ctx context.Context, mid MeetupID, key Key, value Value) (<-chan MeetupUserLink, <-chan error) {
	outC := make(chan MeetupUserLink)
	errC := make(chan error)
	go func() {
		keys := make([]string, 0)
		for k, _ := range r.meetupUsers {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			m := r.meetupUsers[k]
			var thistag = m[string(key)]
			var mu = strings.SplitN(k, "+", 2)
			if mu[0] == string(mid) && thistag != string(value) {
				outC <- MeetupUserLink{MeetupID(mu[0]), UserID(mu[1])}
			}
		}
		close(outC)
		close(errC)
	}()

	return outC, errC
}

func (r *mockRepository) FindMeetupUserLinksMatch(ctx context.Context, mid MeetupID, key Key, value Value) (<-chan MeetupUserLink, <-chan error) {
	return nil, nil
}

func (r *mockRepository) GetPollResult(context.Context, PollID, Key) (PollResult, error) {
	return nil, nil
}
func (r *mockRepository) SetPollUserLinkProperties(context.Context, PollUserLink, KVMap) error {
	return nil
}
func (r *mockRepository) FindPollUserLinksNotMatch(context.Context, PollID, Key, Value) (<-chan PollUserLink, <-chan error) {
	return nil, nil
}
