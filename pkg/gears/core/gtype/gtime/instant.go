package gtime

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

var (
	// Zero represents the zero time.Time value in Go.
	Zero = NewInstant(0).Time()
	// Max represents the max time.Time value in Go.
	Max = Zero.Add(math.MaxInt64)
	// Min represents the min time.Time value in Go.
	Min = Zero.Add(math.MinInt64)
)

func NewInstant(nanos int64) Instant {
	return Instant{nanos: nanos}
}

func NewInstantFromTime(t time.Time) Instant {
	return Instant{nanos: t.UnixNano()}
}

func Now() Instant {
	return Instant{nanos: time.Now().UnixNano()}
}

// Instant represents a point in time.
type Instant struct {
	// The actual time as a long, e.g. nanoseconds since Epoch.
	nanos int64
}

func (i Instant) String() string {
	return fmt.Sprintf("nanos: [%d]", i.nanos)
}

func (i Instant) Nanos() int64 {
	return i.nanos
}

func (i Instant) Time() time.Time {
	return time.Unix(i.nanos/int64(time.Second), i.nanos%int64(time.Second))
}

func (i Instant) Equals(x Instant) bool {
	return i.nanos == x.nanos
}

func (i Instant) Before(x Instant) bool {
	if i.nanos < x.nanos {
		return true
	}
	return false
}

func (i Instant) After(x Instant) bool {
	if i.nanos > x.nanos {
		return true
	}
	return false
}

func (i Instant) Add(d time.Duration) Instant {
	return Instant{nanos: i.nanos + d.Nanoseconds()}
}

func (i Instant) Truncate(d time.Duration) Instant {
	return Instant{nanos: i.nanos / d.Nanoseconds() * d.Nanoseconds()}
}

func (i Instant) Sub(x Instant) time.Duration {
	return time.Duration(i.nanos - x.nanos)
}

func (i Instant) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.nanos)
}

func (i *Instant) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &i.nanos); err != nil {
		return err
	}
	return nil
}
