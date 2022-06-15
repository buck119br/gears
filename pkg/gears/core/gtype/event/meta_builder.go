package event

import (
	"time"

	"github.com/google/uuid"
)

func NewMetaBuilder() MetaBuilder {
	mb := new(metaBuilder)
	return mb
}

type MetaBuilder interface {
	Build() Meta

	WithEventTime(time.Time) MetaBuilder
	WithId(string) MetaBuilder
	WithEventType(string) MetaBuilder
	WithSender(string) MetaBuilder
}

type metaBuilder struct {
	eventTime time.Time
	id        string
	eventType string
	sender    string
}

func (b *metaBuilder) Build() Meta {
	if b.eventTime.IsZero() {
		b.eventTime = time.Now()
	}

	if len(b.id) == 0 {
		b.id = uuid.New().String()
	}

	m := &MetaT{
		EventTimeF: b.eventTime,
		IdF:        b.id,
		EventTypeF: b.eventType,
		SenderF:    b.sender,
	}

	return m
}

func (b *metaBuilder) WithEventTime(t time.Time) MetaBuilder {
	b.eventTime = t
	return b
}

func (b *metaBuilder) WithId(id string) MetaBuilder {
	b.id = id
	return b
}

func (b *metaBuilder) WithEventType(t string) MetaBuilder {
	b.eventType = t
	return b
}

func (b *metaBuilder) WithSender(s string) MetaBuilder {
	b.sender = s
	return b
}
