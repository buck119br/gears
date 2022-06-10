package common

import (
	"fmt"
)

type Key interface {
	fmt.Stringer

	Value() []byte
}

func NewKey(v []byte) Key {
	k := &key{
		v: v,
	}

	return k
}

type key struct {
	v []byte
}

func (k *key) String() string {
	return AnyToString(k.v)
}

func (k *key) Value() []byte {
	return k.v
}
