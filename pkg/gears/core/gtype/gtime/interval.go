package gtime

import (
	"fmt"
	"strconv"
	"strings"
)

// Interval represents a half open interval in time.
type Interval struct {
	// The start point of the interval, inclusive
	s Instant
	// The end point of the interval, optional
	e Instant
}

func NewInterval(startPoint, endPoint Instant) (Interval, error) {
	if endPoint.Nanos() < startPoint.Nanos() {
		return Interval{}, ErrInvalidEndPoint
	}

	return Interval{s: startPoint, e: endPoint}, nil
}

func ParseInterval(s string) (Interval, error) {
	slice := strings.Split(s, "-")
	if len(slice) != 2 {
		return Interval{s: Min, e: Max}, fmt.Errorf("invalid interval format: [%s]", s)
	}
	sNanos, err := strconv.ParseInt(slice[0], 10, 64)
	if err != nil {
		return Interval{s: Min, e: Max}, fmt.Errorf("parse start: [%s] error: [%v]", slice[0], err)
	}
	eNanos, err := strconv.ParseInt(slice[1], 10, 64)
	if err != nil {
		return Interval{s: Min, e: Max}, fmt.Errorf("parse end: [%s] error: [%v]", slice[1], err)
	}
	return Interval{s: NewInstant(sNanos), e: NewInstant(eNanos)}, nil
}

func (i Interval) StartPoint() Instant {
	return i.s
}

func (i Interval) EndPoint() Instant {
	return i.e
}

func (i Interval) String() string {
	return fmt.Sprintf("%d-%d", i.s.nanos, i.e.nanos)
}

func (i Interval) IsBounded() bool {
	if i.e.Equals(Max) {
		return false
	}

	return true
}

func (i Interval) Length() int64 {
	if i.e.Equals(Max) {
		return InfLength
	}

	return i.e.Nanos() - i.s.Nanos()
}

func (i Interval) Equals(x Interval) bool {
	return i.s.Equals(x.s) && i.e.Equals(x.e)
}
