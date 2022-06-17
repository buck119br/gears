package common

import (
	"fmt"
)

type Key interface {
	fmt.Stringer

	Value() string
}

func NewKey(v string) Key {
	k := key{
		v: v,
	}

	return k
}

type key struct {
	v string
}

func (k key) String() string {
	return k.v
}

func (k key) Value() string {
	return k.v
}
