package kv

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type KV interface {
	element.Element
	fmt.Stringer

	Value() common.Value

	WithTimestamp(time.Time) KV
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
	i time.Time
}

func (x *kv) Type() element.Type {
	return element.KV
}

func (x *kv) Key() common.Key {
	return x.k
}

func (x *kv) Timestamp() time.Time {
	return x.i
}

func (x *kv) String() string {
	return fmt.Sprintf("key: [%s], value: [%s]", x.k, x.v)
}

func (x *kv) Value() common.Value {
	return x.v
}

func (x *kv) WithTimestamp(i time.Time) KV {
	x.i = i
	return x
}
