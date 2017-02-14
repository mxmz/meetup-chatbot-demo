package gaeapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"telegram-meetup/app"
	"telegram-meetup/env"
	. "telegram-meetup/types"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine/urlfetch"
)

func init() {

}

type messageQueue struct {
	env        Env
	omh        OutboundChatMessageHandler
	imh        InboundChatMessageHandler
	wmh        WebhookMessageHandler
	jrh        JobRequesteHandler
	rawMessage struct {
		Type string
		Data []byte
	}
}

func (mq *messageQueue) MakeEnv() Env { return mq.env }

func (mq *messageQueue) RegisterOutboundChatMessagehandler(h OutboundChatMessageHandler) error {
	mq.omh = h
	return nil
}
func (mq *messageQueue) RegisterInboundChatMessagehandler(h InboundChatMessageHandler) error {
	mq.imh = h
	return nil
}
func (mq *messageQueue) RegisterJobRequesteHandler(h JobRequesteHandler) error {
	mq.jrh = h
	return nil
}
func (mq *messageQueue) RegisterWebhookMessageHandler(h WebhookMessageHandler) error {
	mq.wmh = h
	return nil
}
func (mq *messageQueue) MakeRepository() Repository {
	return NewUserStore()
}
func (mq *messageQueue) MakeMeetupService() MeetupService {
	var mc, _ = NewMeetupCollaborators(mq.env)
	return mc
}
func (mq *messageQueue) MakeMeetupAuthorizer() MeetupAuthorizer {
	var _, ma = NewMeetupCollaborators(mq.env)
	return ma
}
func (mq *messageQueue) MakeHttpClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}

func NewMessageQueue() *messageQueue {
	var env = env.LoadEnv()
	return &messageQueue{env: env}
}

const MessageQueueHandler = "/message-queue-handler"
const OutboundMessageType = "outbound_message"
const InboundMessageType = "inbound_message"
const JobRequestType = "job_request"
const WebhookMessageType = "webhook_message"

const InboundMessageQueueName = "inbound-messages"
const OutboundMessageQueueName = "outbound-messages"
const JobRequestQueueName = "jobs"
const WebhookMessageName = "webhook-messages"

func (e *messageQueue) HandleWebhookMessage(ctx context.Context, m *WebhookMessage) error {
	taskURL := MessageQueueHandler + "?type=" + WebhookMessageType
	t := NewPOSTTask(taskURL, m)
	_, err := taskqueue.Add(ctx, t, WebhookMessageName)
	return err

}

func (e *messageQueue) HandleOutboundChatMessage(ctx context.Context, m *OutboundChatMessage) error {
	taskURL := MessageQueueHandler + "?type=" + OutboundMessageType
	t := NewPOSTTask(taskURL, m)
	_, err := taskqueue.Add(ctx, t, OutboundMessageQueueName)
	return err

}

func (e *messageQueue) HandleInboundChatMessage(ctx context.Context, m *InboundChatMessage) error {
	taskURL := MessageQueueHandler + "?type=" + InboundMessageType
	t := NewPOSTTask(taskURL, m)
	_, err := taskqueue.Add(ctx, t, InboundMessageQueueName)
	return err

}

func (e *messageQueue) HandleJobRequest(ctx context.Context, m *JobRequest) error {
	taskURL := MessageQueueHandler + "?type=" + JobRequestType
	t := NewPOSTTask(taskURL, m)
	_, err := taskqueue.Add(ctx, t, JobRequestQueueName)
	return err
}

func NewPOSTTask(path string, data interface{}) *taskqueue.Task {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	var json, _ = json.Marshal(data)
	return &taskqueue.Task{
		Path:    path,
		Payload: json,
		Header:  h,
		Method:  "POST",
	}
}

func (mq *messageQueue) UnmarshalInboundChatMessage(bytes []byte) (*InboundChatMessage, error) {
	m := InboundChatMessage{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (mq *messageQueue) UnmarshalOutboundChatMessage(bytes []byte) (*OutboundChatMessage, error) {
	m := OutboundChatMessage{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (mq *messageQueue) UnmarshalJobRequest(bytes []byte) (*JobRequest, error) {
	m := JobRequest{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (mq *messageQueue) UnmarshalWebhookMessage(bytes []byte) (*WebhookMessage, error) {
	m := WebhookMessage{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

var errMustDefer = errors.New("message should be deferred")

func (mq *messageQueue) dispatch(ctx context.Context) error {
	msgType := mq.rawMessage.Type
	body := mq.rawMessage.Data
	switch msgType {
	case InboundMessageType:
		{
			m, err := mq.UnmarshalInboundChatMessage(body)
			if err != nil {
				return err
			}
			err = mq.imh.HandleInboundChatMessage(ctx, m)
			if err != nil {
				return err
			}
			return nil
		}
	case JobRequestType:
		{
			m, err := mq.UnmarshalJobRequest(body)
			if err != nil {
				return err
			}
			now := time.Now()

			if m.Scheduled.After(now) {
				return errMustDefer
			}
			if m.Droppable.Before(now) {
				glog.Debugf(ctx, "dropped: %v", m)
				return nil
			}

			err = mq.jrh.HandleJobRequest(ctx, m)
			if err != nil {
				return err
			}
			return nil
		}
	case OutboundMessageType:
		{
			m, err := mq.UnmarshalOutboundChatMessage(body)
			if err != nil {
				return err
			}
			err = mq.omh.HandleOutboundChatMessage(ctx, m)
			if err != nil {
				return err
			}
			return nil
		}
	case WebhookMessageType:
		{
			m, err := mq.UnmarshalWebhookMessage(body)
			if err != nil {
				return err
			}
			err = mq.wmh.HandleWebhookMessage(ctx, m)
			if err != nil {
				return err
			}
			return nil
		}
	default:
		{
			return errors.New("unexpected message type")
		}
	}
	//glog.Debugf(ctx, "err: %s", "should never be here")
	//return nil
}

func SetupMessageQueueHandler(env Env) {

	http.HandleFunc(MessageQueueHandler,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			ctx := appengine.NewContext(r)
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(461)
				glog.Debugf(ctx, "err: %v", err)
				fmt.Fprint(w, err)
			} else {
				var mq = NewMessageQueue()
				app.RegisterChatBotApp(mq)
				app.RegisterInboundWorker(mq)
				app.RegisterOutboundWorker(mq)

				q := r.URL.Query()
				msgType := q.Get("type")
				mq.rawMessage = struct {
					Type string
					Data []byte
				}{msgType, body}
				err := mq.dispatch(ctx)
				if err != nil {
					w.WriteHeader(461)
					glog.Debugf(ctx, "err: %v", err)
					fmt.Fprint(w, err)
				} else {
					w.WriteHeader(200)
					fmt.Fprint(w, "ok")
				}

			}
		})
}

func must(f func() error) {
	err := f()
	if err != nil {
		panic(err)
	}
}
