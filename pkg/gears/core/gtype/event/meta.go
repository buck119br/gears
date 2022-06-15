package event

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

// NewMeta returns an uninitialized instance of Meta.
// MetaBuilder is recommended for creating an Event.
func NewMeta() Meta {
	return new(MetaT)
}

type Meta interface {
	fmt.Stringer

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

func (m *MetaT) String() string {
	return fmt.Sprintf("event time: [%s], id: [%s], event type: [%s], sender: [%s]",
		common.AnyToString(m.EventTimeF), m.IdF, m.EventTypeF, m.SenderF)
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
