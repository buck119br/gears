package gtime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Interval represents a half open interval in time.
type Interval struct {
	// The start point of the interval, inclusive
	s time.Time
	// The end point of the interval, optional
	e time.Time
}

func NewInterval(startPoint, endPoint time.Time) (Interval, error) {
	if endPoint.Before(startPoint) {
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
	return Interval{s: NewInstant(sNanos).Time(), e: NewInstant(eNanos).Time()}, nil
}

func (i Interval) String() string {
	return fmt.Sprintf("%d-%d", i.s.UnixNano(), i.e.UnixNano())
}

func (i Interval) StartPoint() time.Time {
	return i.s
}

func (i Interval) EndPoint() time.Time {
	return i.e
}

func (i Interval) IsBounded() bool {
	if i.e.Equal(Max) {
		return false
	}

	return true
}

func (i Interval) Length() int64 {
	if i.e.Equal(Max) {
		return InfLength
	}

	return i.e.UnixNano() - i.s.UnixNano()
}

func (i Interval) Equals(x Interval) bool {
	return i.s.Equal(x.s) && i.e.Equal(x.e)
}
