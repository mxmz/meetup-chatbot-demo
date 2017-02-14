package types

import "time"

type UserID string
type MeetupID string
type MeetupUserLink struct {
	MeetupID
	UserID
}
type Key string
type Value string
type KVMap map[Key]Value

type PollID string
type PollResult map[string]int
type PollUserLink struct {
	PollID
	UserID
}

type WebhookMessage struct {
	Type string
	Data []byte
}

type Location struct{ Lat, Lng float64 }
type Button struct{ Text, Command string }
type File struct {
	Name string
	Data []byte
}
type OutboundChatMessage struct {
	RecipientID string    `json:"RID"`
	Message     string    `json:"Msg"`
	Map         *Location `json:"Map"`
	Buttons     []Button  `json:"But"`
	Attachment  *File     `json:"Fil"`
}

type InboundChatMessage struct {
	SenderID   string   `msg:"0"`
	SenderName string   `msg:"1"`
	Message    string   `msg:"2"`
	Command    []string `msg:"3"`
}

type JobRequest struct {
	Cookie    []byte    `msg:"0"`
	Command   []string  `msg:"1"`
	Scheduled time.Time `msg:"2"`
	Droppable time.Time `msg:"2"`
}

type MeetupGroupID string

type MeetupAuthTokens []byte
type MeetupAuthState []byte

type MeetupAuthCompletion struct {
	State  MeetupAuthState
	Tokens MeetupAuthTokens
}

type MeetupGroup struct {
	Name    string
	Urlname string
	ID      int
}

type Meetup struct {
	ETag  string
	Title string
	Where string
	When  time.Time
	Map   *Location
}

type URL string
