package mansbridge

import (
	"time"
)

type WireMessage struct {
	Entries     Entries
	Source      string
	CurrentTime time.Time
}

type WireService interface {
	SendNews(correspondent string, entries []Entry)
	GetNews() <-chan WireMessage
	GetAddress() string
}
