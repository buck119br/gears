package row

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

type Column interface {
	fmt.Stringer

	Key() common.Key
	Value() common.Value
	Type() common.Type
	IsPK() bool

	WithValue(common.Value) Column
}

type column struct {
	k    common.Key
	v    common.Value
	t    common.Type
	isPK bool
}

func NewColumn(k common.Key, v common.Value, t common.Type, isPK bool) Column {
	c := new(column)
	c.k = k
	c.v = v
	c.t = t
	c.isPK = isPK

	return c
}

func (c *column) String() string {
	return fmt.Sprintf("key: [%s], type: [%s], value: [%s], isPK: [%t]", c.k, c.t, common.AnyToString(c.v.As(c.t)), c.isPK)
}

func (c *column) Key() common.Key {
	return c.k
}

func (c *column) Value() common.Value {
	return c.v
}

func (c *column) Type() common.Type {
	return c.t
}

func (c *column) IsPK() bool {
	return c.isPK
}

func (c *column) WithValue(v common.Value) Column {
	c.v = v
	return c
}
