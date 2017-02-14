package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"telegram-meetup/env"
	"telegram-meetup/host/embed"
	"telegram-meetup/meetup"
	"telegram-meetup/tskbroker"
	. "telegram-meetup/types"
	types "telegram-meetup/types"
	"time"

	"github.com/pborman/uuid"

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
	E = env.LoadEnv()
}

type appHost struct {
	mq         *tskbroker.MqBrokerMap
	omh        OutboundChatMessageHandler
	imh        InboundChatMessageHandler
	wmh        WebhookMessageHandler
	jrh        JobRequesteHandler
	queueNames map[string]struct{}
}

func (d *appHost) injectMessage(ctx context.Context, typ string, payload interface{}, delay time.Time) error {
	msg := &tskbroker.MqMessage{
		ID:      uuid.New(),
		Type:    typ,
		Delayed: delay,
		Payload: payload,
	}
	log.Println("inject", typ, msg)
	return d.mq.Inject(ctx, typ, msg)
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

func NewAppHost(mq *tskbroker.MqBrokerMap) *appHost {
	return &appHost{mq: mq, queueNames: make(map[string]struct{})}
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
	d.mq.Listen(ctx, func(m *tskbroker.MqMessage) error {
		msgType := m.Type
		payload := m.Payload
		switch msgType {
		case InboundMessageQueueName:
			{
				m, _ := payload.(*InboundChatMessage)
				err := d.imh.HandleInboundChatMessage(ctx, m)
				if err != nil {
					return err
				}
				return nil
			}
		case JobRequestQueueName:
			{
				m, _ := payload.(*JobRequest)
				now := time.Now()
				if m.Droppable.Before(now) {
					log.Println("dropped:", m)
					return nil
				}
				err := d.jrh.HandleJobRequest(ctx, m)
				if err != nil {
					return err
				}
				return nil
			}
		case OutboundMessageQueueName:
			{
				m, _ := payload.(*OutboundChatMessage)

				err := d.omh.HandleOutboundChatMessage(ctx, m)
				if err != nil {
					return err
				}
				return nil
			}
		case WebhookMessageName:
			{
				m, _ := payload.(*WebhookMessage)
				err := d.wmh.HandleWebhookMessage(ctx, m)
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
