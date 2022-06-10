package common

import (
	"fmt"
	"strconv"
	"time"
)

type Value interface {
	fmt.Stringer

	Value() []byte
	As(Type) interface{}
}

func NewValue(v []byte) Value {
	val := new(value)
	val.v = v

	return val
}

type value struct {
	v []byte
}

func (v *value) String() string {
	return bytesToString(v.v)
}

func (v *value) Value() []byte {
	return v.v
}

func (v *value) As(t Type) interface{} {
	var x interface{}
	var err error
	switch t {
	case Bool:
		x, err = strconv.ParseBool(bytesToString(v.v))
	case Integer:
		x, err = strconv.ParseInt(bytesToString(v.v), 10, 64)
	case Float:
		x, err = strconv.ParseFloat(bytesToString(v.v), 64)
	case String:
		x, err = bytesToString(v.v), nil
	case Binary:
		x, err = v.v, nil
	case Timestamp:
		x, err = time.Parse(time.RFC3339, bytesToString(v.v))
	default:
		x, err = nil, fmt.Errorf("invalid column type: [%d]", t)
	}

	if err != nil {
		panic(fmt.Errorf("cannot parse value: [%s] as: [%s]", bytesToString(v.v), t))
	}

	return x
}
