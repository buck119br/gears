package kv

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
	"github.com/buck119br/gears/pkg/gears/core/gtype/gtime"
)

type KV interface {
	element.Element

	Value() common.Value

	WithTimestamp(gtime.Instant) KV
}

func NewKV(k common.Key, v common.Value) KV {
	x := new(kv)
	x.k = k
	x.v = v

	return x
}

type kv struct {
	k common.Key
	v common.Value
	i gtime.Instant
}

func (x *kv) Type() element.Type {
	return element.KV
}

func (x *kv) Key() common.Key {
	return x.k
}

func (x *kv) Timestamp() gtime.Instant {
	return x.i
}

func (x *kv) String() string {
	return fmt.Sprintf("key: [%s], value: [%s]", x.k, x.v)
}

func (x *kv) Value() common.Value {
	return x.v
}

func (x *kv) WithTimestamp(i gtime.Instant) KV {
	x.i = i
	return x
}
