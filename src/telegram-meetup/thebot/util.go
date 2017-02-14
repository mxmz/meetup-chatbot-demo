package thebot

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	. "telegram-meetup/types"
	"time"
)

func md5hex(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func eventInProgress(e *Meetup) bool {
	now := time.Now()
	rv := e.When.Before(now) && e.When.Add(4*time.Hour).After(now)
	log.Println("eventInProgress", e.When, e.When.Add(4*time.Hour), now, rv)
	return rv
}
