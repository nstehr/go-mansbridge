package wire

import (
	"github.com/nstehr/go-mansbridge/collections"
	"time"
)

type WireMessage struct {
	Entries     []collections.Entry
	Source      string
	CurrentTime time.Time
}

type WireService interface {
	SendNews(correspondent string, entries []collections.Entry)
	GetNews() <-chan WireMessage
	GetAddress() string
}
