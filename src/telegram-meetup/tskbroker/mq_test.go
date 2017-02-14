package tskbroker

import (
	"context"
	"errors"
	"log"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

var _ = log.Println

func TestSimple(t *testing.T) {
	RegisterTestingT(t)

	mqbs := NewMqBrokerMap(context.TODO(), 1, 10)

	var wg sync.WaitGroup
	wg.Add(1)

	mqb := mqbs.GetBrokerByName("foobar")

	go func() {

		wg.Done()
	}()
	wg.Add(1)
	go func() {

		for i := 0; i < 100; i++ {
			msg := &MqMessage{
				ID:      "asdakjdhakjdh",
				Type:    "Boh",
				Delayed: time.Now().Add(time.Duration(i/10) * time.Second),
				Payload: 290,
			}
			mqbs.Inject(context.TODO(), "foobar", msg)
		}
		log.Println("produced")
		defer wg.Done()
	}()
	var msgs [3]MqMessage

	wg.Add(1)
	go func() {
		//defer mqb.Close()
		i := 0
		mqbs.Listen(mqb.Cancel(), func(m *MqMessage) error {
			defer func() {
				i++
			}()
			log.Println(m)
			if i < 3 {
				msgs[i] = *m
				return errors.New("error")
			} else {
				return nil
			}
		}, "foobar")
		log.Println("consumed")
		wg.Done()
	}()
	<-time.After(30 * time.Second)
	mqbs.Shutdown()
	wg.Wait()
	Expect(msgs[0].ID).To(BeEquivalentTo("asdakjdhakjdh"))

	//<-c1.MsgC

	//<-c2.MsgC
	//<-c3.inC
}
