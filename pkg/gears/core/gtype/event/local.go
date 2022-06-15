package event

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

func NewLocal() Local {
	return new(local)
}

type Local interface {
	fmt.Stringer

	ProcessingTime() time.Time
	IngestionTime() time.Time
	Key() common.Key
	// Destination represents the destination of output I/O (such as Kafka topic etc.), which to write to.
	Destination() string

	WithProcessingTime(time.Time) Local
	WithIngestionTime(time.Time) Local
	WithKey(common.Key) Local
	WithDestination(string) Local
}

type local struct {
	ProcessingTimeF time.Time `json:"processing_time"`
	IngestionTimeF  time.Time `json:"ingestion_time"`

	KeyF         common.Key `json:"key"`
	DestinationF string     `json:"destination"`
}

func (l *local) String() string {
	return fmt.Sprintf("processing time: [%s], ingestion time: [%s], key: [%s], destination: [%s]",
		common.AnyToString(l.ProcessingTimeF), common.AnyToString(l.IngestionTimeF), l.KeyF, l.DestinationF)
}

func (l *local) ProcessingTime() time.Time {
	return l.ProcessingTimeF
}

func (l *local) IngestionTime() time.Time {
	return l.IngestionTimeF
}

func (l *local) Key() common.Key {
	return l.KeyF
}

func (l *local) Destination() string {
	return l.DestinationF
}

func (l *local) WithProcessingTime(t time.Time) Local {
	l.ProcessingTimeF = t
	return l
}

func (l *local) WithIngestionTime(t time.Time) Local {
	l.IngestionTimeF = t
	return l
}

func (l *local) WithKey(k common.Key) Local {
	l.KeyF = k
	return l
}

func (l *local) WithDestination(t string) Local {
	l.DestinationF = t
	return l
}
