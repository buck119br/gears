package event

import (
	"time"
)

// NewMeta returns an uninitialized instance of Meta.
// MetaBuilder is recommended for creating an Event.
func NewMeta() Meta {
	return new(MetaT)
}

type Meta interface {
	EventTime() time.Time
	Id() string
	EventType() string
	Sender() string
}

type MetaT struct {
	EventTimeF time.Time `json:"event_time" avro:"event_time"`
	IdF        string    `json:"id" avro:"id"`
	EventTypeF string    `json:"event_type" avro:"event_type"`
	SenderF    string    `json:"sender" avro:"sender"`
}

func (m *MetaT) EventTime() time.Time {
	return m.EventTimeF
}

func (m *MetaT) Id() string {
	return m.IdF
}

func (m *MetaT) EventType() string {
	return m.EventTypeF
}

func (m *MetaT) Sender() string {
	return m.SenderF
}
