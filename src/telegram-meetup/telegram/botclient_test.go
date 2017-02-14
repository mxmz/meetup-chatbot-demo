package telegram_test

import (
	"net/http"
	"os"
	"telegram-meetup/telegram"
	. "telegram-meetup/types"
	"testing"

	"golang.org/x/net/context"

	. "github.com/onsi/gomega"
)

type depMock struct {
}

func defaultClient(context.Context) *http.Client {
	return http.DefaultClient
}

func (*depMock) Get(k string) string {
	return os.Getenv(k)
}
func (*depMock) HandleInboundChatMessage(context.Context, *InboundChatMessage) error {
	return nil
}

func TestNewBot(t *testing.T) {
	RegisterTestingT(t)
	os.Setenv("TELEGRAM_TOKEN", "0")
	b, err := telegram.NewBot(&depMock{}, &depMock{}, defaultClient)
	Expect(err.Error()).To(BeEquivalentTo("undefined or invalid Telegram token: 0"))
	Expect(b).To(BeNil())

}
