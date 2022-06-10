package gtime

import (
	"fmt"
	"time"
)

type TTL interface {
	fmt.Stringer

	Duration() time.Duration
}

func NewTTL(d time.Duration) TTL {
	t := new(ttl)
	t.d = d

	return t
}

type ttl struct {
	d time.Duration
}

func (t *ttl) String() string {
	return fmt.Sprintf("duration: [%s]", t.d)
}

func (t *ttl) Duration() time.Duration {
	return t.d
}
