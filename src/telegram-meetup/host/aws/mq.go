package aws

import (
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"golang.org/x/net/context"
)

type MqMessage struct {
	ID      string
	Type    string
	Delayed time.Time
	Payload string
}

type MqMap struct {
	sess *session.Session
	urls map[string]string
	mtx  sync.Mutex
}

func NewMqMap(urls map[string]string) *MqMap {
	sess := session.New(&aws.Config{
		Region:     aws.String("us-west-2"),
		DisableSSL: aws.Bool(true),
	})

	return &MqMap{sess: sess, urls: urls}
}

func (mq *MqMap) Inject(ctx context.Context, queue string, m *MqMessage, isFifo bool) error {
	url := mq.urls[queue]
	log.Println(m.Delayed)
	log.Println(time.Now())
	log.Println(int64(m.Delayed.Sub(time.Now())))
	messageGroupId := aws.String(m.Type)
	messageDeduplicationId := aws.String(m.ID)
	delaySeconds := aws.Int64(int64(m.Delayed.Sub(time.Now()) / time.Second))
	if isFifo {
		delaySeconds = nil
	} else {
		messageDeduplicationId = nil
		messageGroupId = nil
	}

	svc := sqs.New(mq.sess)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:               aws.String(url),
		MessageBody:            aws.String(m.Payload),
		MessageGroupId:         messageGroupId,
		MessageDeduplicationId: messageDeduplicationId,
		DelaySeconds:           delaySeconds,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (mq *MqMap) Listen(ctx context.Context, f func(m *MqMessage) error, queues ...string) {
	type ReceiveMessagesWithName struct {
		qName string
		msgs  *sqs.ReceiveMessageOutput
	}

	read := make(chan ReceiveMessagesWithName)
	ctx, cancel := context.WithCancel(ctx)
	urls := map[string]string{}
	for _, n := range queues {
		url, ok := mq.urls[n]
		if !ok {
			panic("queue not define: " + n)
		}
		urls[n] = url
	}
	for qName, _ := range urls {
		go func(qName string) {
			defer log.Println("quit listener")
			defer cancel()
			url := urls[qName]
			svc := sqs.New(mq.sess)
			for {
				log.Println("waiting for for", url, "...")
				recParams := &sqs.ReceiveMessageInput{
					QueueUrl:            aws.String(url),
					MaxNumberOfMessages: aws.Int64(3),
					VisibilityTimeout:   aws.Int64(30),
					WaitTimeSeconds:     aws.Int64(10),
				}
				recResp, err := svc.ReceiveMessage(recParams)
				if err != nil {
					log.Println(err)
				}
				if len(recResp.Messages) > 0 {
					log.Println("emit msg...", qName)
					select {
					case <-ctx.Done():
						{
							return
						}
					case read <- ReceiveMessagesWithName{qName, recResp}:
					}
					log.Println("emitted.")
				} else {
					log.Println("check ctx...")
					select {
					case <-ctx.Done():
						{
							return
						}
					default:
					}
				}
			}
		}(qName)
	}
	for {
		svc := sqs.New(mq.sess)
		log.Println("collecting msg...")
		select {
		case <-ctx.Done():
			{
				return
			}

		case recResp := <-read:
			{
				for _, awMsg := range recResp.msgs.Messages {
					_ = awMsg
					body := awMsg.Body
					msg := MqMessage{
						ID:      *awMsg.MessageId,
						Type:    recResp.qName,
						Payload: *body,
					}
					log.Println(msg.ID, msg.Type, len(msg.Payload))
					err := f(&msg)
					if err == nil {
						delParams := &sqs.DeleteMessageInput{
							QueueUrl:      aws.String(urls[recResp.qName]),
							ReceiptHandle: awMsg.ReceiptHandle,
						}
						_, err := svc.DeleteMessage(delParams)
						if err != nil {
							log.Println(err)
						}
					} else {
						log.Println("not deleting:", awMsg.MessageId, err)
					}

					select {
					case <-ctx.Done():
						{
							return
						}
					default:
					}
				}

			}
		}
	}

}
