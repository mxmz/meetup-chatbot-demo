package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"telegram-meetup/env"
	"telegram-meetup/host/embed"
	"telegram-meetup/meetup"
	. "telegram-meetup/types"
	types "telegram-meetup/types"
	"time"

	"github.com/pborman/uuid"

	"telegram-meetup/host/aws"

	"golang.org/x/net/context"
)

const InboundMessageQueueName = "inbound-messages"
const OutboundMessageQueueName = "outbound-messages"
const JobRequestQueueName = "jobs"
const WebhookMessageName = "webhook-messages"

var dbFile = "./repo.db"
var db *sql.DB
var E types.Env

func init() {

	_, err := os.Lstat(dbFile)
	log.Println(err)
	if err != nil && os.IsNotExist(err) {
		log.Println("creating")
		db, err = sql.Open("sqlite3", dbFile)
		if err != nil {
			panic(err)
		}
		err = embed.DBCreate(db)
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("opening")
		db, err = sql.Open("sqlite3", dbFile)
		if err != nil {
			panic(err)
		}
	}
	env.DevMode = true
	env := env.LoadEnv()
	env.Export()
	E = env
}

type appHost struct {
	mq         *aws.MqMap
	omh        OutboundChatMessageHandler
	imh        InboundChatMessageHandler
	wmh        WebhookMessageHandler
	jrh        JobRequesteHandler
	queueNames map[string]struct{}
}

func (d *appHost) injectMessage(ctx context.Context, typ string, payload interface{}, delay time.Time) error {
	json, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	msg := &aws.MqMessage{
		ID:      uuid.New(),
		Type:    typ,
		Delayed: delay,
		Payload: string(json),
	}
	log.Println("inject", typ, msg)
	return d.mq.Inject(ctx, typ, msg, typ != JobRequestQueueName)
}

func (d *appHost) HandleInboundChatMessage(ctx context.Context, m *types.InboundChatMessage) error {
	return d.injectMessage(ctx, InboundMessageQueueName, m, time.Now())
}

func (d *appHost) HandleOutboundChatMessage(ctx context.Context, m *types.OutboundChatMessage) error {
	return d.injectMessage(ctx, OutboundMessageQueueName, m, time.Now())
}

func (d *appHost) HandleJobRequest(ctx context.Context, m *types.JobRequest) error {
	return d.injectMessage(ctx, JobRequestQueueName, m, m.Scheduled)
}

func (d *appHost) RegisterOutboundChatMessagehandler(h OutboundChatMessageHandler) error {
	d.omh = h
	d.queueNames[OutboundMessageQueueName] = struct{}{}
	return nil
}
func (d *appHost) RegisterInboundChatMessagehandler(h InboundChatMessageHandler) error {
	d.imh = h
	d.queueNames[InboundMessageQueueName] = struct{}{}
	return nil
}
func (d *appHost) RegisterJobRequesteHandler(h JobRequesteHandler) error {
	d.jrh = h
	d.queueNames[JobRequestQueueName] = struct{}{}
	return nil
}
func (d *appHost) RegisterWebhookMessageHandler(h WebhookMessageHandler) error {
	d.wmh = h
	d.queueNames[WebhookMessageName] = struct{}{}
	return nil
}
func (d *appHost) MakeRepository() Repository {
	repo, _ := embed.NewRepo(E, db)
	return repo
}

func httpDo(ctx context.Context, r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r.WithContext(ctx))
}

func (d *appHost) MakeMeetupService() MeetupService {
	return meetup.NewClient(E, httpDo)
	//var mc, _ = NewMeetupCollaborators(E)
	//return mc
}
func (d *appHost) MakeMeetupAuthorizer() MeetupAuthorizer {
	return nil
	//var _, ma = NewMeetupCollaborators(E)
	//return ma
}
func (d *appHost) MakeHttpClient(ctx context.Context) *http.Client {
	return http.DefaultClient
}
func (d *appHost) MakeEnv() Env { return E }

func NewAppHost() *appHost {
	queues := []string{
		OutboundMessageQueueName,
		InboundMessageQueueName,
		WebhookMessageName,
		JobRequestQueueName,
	}
	urls := map[string]string{}
	for _, n := range queues {
		if n == JobRequestQueueName {
			urls[n] = E.Get("AWS_SQS_BASE_URL") + "/" + n
		} else {
			urls[n] = E.Get("AWS_SQS_BASE_URL") + "/" + n + ".fifo"
		}
	}
	return &appHost{mq: aws.NewMqMap(urls), queueNames: map[string]struct{}{}}
}

func (d *appHost) handleWebhookMessage(ctx context.Context, m *types.WebhookMessage) error {
	return d.injectMessage(ctx, WebhookMessageName, m, time.Now())
}

func (d *appHost) Run(ctx context.Context) {

	queueNames := make([]string, 0)
	for k, _ := range d.queueNames {
		queueNames = append(queueNames, k)
	}
	log.Println("listening ", queueNames)
	d.mq.Listen(ctx, func(m *aws.MqMessage) error {
		msgType := m.Type
		payload := m.Payload
		switch msgType {
		case InboundMessageQueueName:
			{
				var m InboundChatMessage
				json.Unmarshal([]byte(payload), &m)
				err := d.imh.HandleInboundChatMessage(ctx, &m)
				if err != nil {
					return err
				}
				return nil
			}
		case JobRequestQueueName:
			{
				var m JobRequest
				json.Unmarshal([]byte(payload), &m)
				now := time.Now()
				if m.Droppable.Before(now) {
					log.Println("dropped:", m)
					return nil
				}
				err := d.jrh.HandleJobRequest(ctx, &m)
				if err != nil {
					return err
				}
				return nil
			}
		case OutboundMessageQueueName:
			{
				var m OutboundChatMessage
				json.Unmarshal([]byte(payload), &m)
				err := d.omh.HandleOutboundChatMessage(ctx, &m)
				if err != nil {
					return err
				}
				return nil
			}
		case WebhookMessageName:
			{
				var m WebhookMessage
				json.Unmarshal([]byte(payload), &m)
				err := d.wmh.HandleWebhookMessage(ctx, &m)
				if err != nil {
					return err
				}
				return nil
			}
		default:
			{
				return errors.New("unexpected message type: " + msgType)
			}
		}
		//glog.Debugf(ctx, "err: %s", "should never be here")
		//return nil

		return nil
	}, queueNames...)
	log.Println("stop listening ")
}
