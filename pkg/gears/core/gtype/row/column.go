package row

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

type Column interface {
	fmt.Stringer

	Name() common.Key
	Value() common.Value
	Type() common.Type
	IsPK() bool

	WithValue(common.Value) Column
}

func NewColumn[V any](n common.Key, v common.Value, t common.Type, isPK bool) Column {
	c := &column{
		n:    n,
		v:    v,
		t:    t,
		isPK: isPK,
	}

	return c
}

type column struct {
	n    common.Key
	v    common.Value
	t    common.Type
	isPK bool
}

func (c *column) String() string {
	return fmt.Sprintf("name: [%s], type: [%s], value: [%s], isPK: [%t]", c.n, c.t, c.v, c.isPK)
}

func (c *column) Name() common.Key {
	return c.n
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
