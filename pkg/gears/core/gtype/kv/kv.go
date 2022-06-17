package kv

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type KV[V any] interface {
	element.Element
	fmt.Stringer

	Value() common.Value

	WithTimestamp(time.Time) KV[V]
}

func NewKV[V any](k common.Key, v common.Value) KV[V] {
	x := &kv[V]{
		k: k,
		v: v,
	}

	return x
}

type kv[V any] struct {
	k common.Key
	v common.Value
	i time.Time
}

func (x *kv[V]) Type() element.Type {
	return element.KV
}

func (x *kv[V]) Key() common.Key {
	return x.k
}

func (x *kv[V]) Timestamp() time.Time {
	return x.i
}

func (x *kv[V]) String() string {
	return fmt.Sprintf("key: [%s], value: [%s]", x.k, x.v)
}

func (x *kv[V]) Value() common.Value {
	return x.v
}

func (x *kv[V]) WithTimestamp(i time.Time) KV[V] {
	x.i = i
	return x
}
