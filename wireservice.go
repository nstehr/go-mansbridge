package mansbridge

import (
	"time"
)

type WireMessage struct {
	Entries     []Entry
	Source      string
	CurrentTime time.Time
}

type WireService interface {
	SendNews(correspondent string, entries []Entry)
	GetNews() <-chan WireMessage
	GetAddress() string
}
