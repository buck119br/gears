package common

import (
	"fmt"
)

type Value interface {
	fmt.Stringer

	Value() any
}

func NewValue(v any) Value {
	val := &value{
		v: v,
	}
	return val
}

type value struct {
	v any
}

func (v *value) String() string {
	return AnyToString(v.v)
}

func (v *value) Value() any {
	return v.v
}
