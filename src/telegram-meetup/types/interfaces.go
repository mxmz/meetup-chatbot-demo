package types

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

type Repository interface {
	SetUserProperties(context.Context, UserID, KVMap) error
	GetUserProperties(context.Context, UserID, ...Key) (KVMap, error)
	SetMeetupUserLinkProperties(context.Context, MeetupUserLink, KVMap) error
	FindMeetupUserLinksNotMatch(context.Context, MeetupID, Key, Value) (<-chan MeetupUserLink, <-chan error)
	FindMeetupUserLinksMatch(context.Context, MeetupID, Key, Value) (<-chan MeetupUserLink, <-chan error)

	GetPollResult(context.Context, PollID, Key) (PollResult, error)
	SetPollUserLinkProperties(context.Context, PollUserLink, KVMap) error
	FindPollUserLinksNotMatch(context.Context, PollID, Key, Value) (<-chan PollUserLink, <-chan error)
}

type InboundChatMessageHandler interface {
	HandleInboundChatMessage(context.Context, *InboundChatMessage) error
}
type JobRequesteHandler interface {
	HandleJobRequest(context.Context, *JobRequest) error
}
type WebhookMessageHandler interface {
	HandleWebhookMessage(context.Context, *WebhookMessage) error
}

type App interface {
	InboundChatMessageHandler
	JobRequesteHandler
	HandleMeetupAuthCompletion(ctx context.Context, mc *MeetupAuthCompletion) error
}

type OutboundChatMessageHandler interface {
	HandleOutboundChatMessage(context.Context, *OutboundChatMessage) error
}

type MeetupService interface {
	GetGroupList(context.Context, MeetupAuthTokens) ([]MeetupGroup, error)
	GetNextMeetup(context.Context, MeetupGroupID, time.Time) (*Meetup, error)
}

var ErrMeetup401 = errors.New("401 Unauthorized")

type MeetupAuthorizer interface {
	MakeOOBAuthStartURL(context.Context, MeetupAuthState) (URL, error)
	RefreshAuthTokens(context.Context, MeetupAuthTokens) (MeetupAuthTokens, error)
}

type Env interface {
	Get(string) string
}
type MessageEnqueuer interface {
	InboundChatMessageHandler
	OutboundChatMessageHandler
	JobRequesteHandler
}

type AppHost interface {
	MessageEnqueuer
	RegisterOutboundChatMessagehandler(OutboundChatMessageHandler) error
	RegisterInboundChatMessagehandler(InboundChatMessageHandler) error
	RegisterJobRequesteHandler(JobRequesteHandler) error
	RegisterWebhookMessageHandler(WebhookMessageHandler) error
	MakeRepository() Repository
	MakeMeetupService() MeetupService
	MakeMeetupAuthorizer() MeetupAuthorizer
	MakeHttpClient(context.Context) *http.Client
	MakeEnv() Env
}
