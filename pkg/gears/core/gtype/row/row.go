package row

import (
	"fmt"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type Row interface {
	element.Element
	fmt.Stringer

	Driver() Driver
	Collection() Collection
	Operation() Operation
	Columns() Columns

	WithKey(common.Key) Row
	WithTimestamp(time.Time) Row

	Check() error
}

func NewRow(d Driver, c Collection, op Operation, cols Columns) Row {
	r := &row{
		d:    d,
		c:    c,
		op:   op,
		cols: cols,
	}

	return r
}

type row struct {
	d    Driver
	c    Collection
	op   Operation
	cols Columns

	key common.Key
	t   time.Time
}

func (r *row) Type() element.Type {
	return element.Row
}

func (r *row) Key() common.Key {
	return r.key
}

func (r *row) Timestamp() time.Time {
	return r.t
}

func (r *row) String() string {
	var dStr, cStr string
	if r.d != nil {
		dStr = r.d.String()
	}
	if r.c != nil {
		cStr = r.c.Name()
	}

	return fmt.Sprintf("driver: [%s], collection: [%s], operation: [%s], columns: [%s], key: [%s], timestamp: [%s]",
		dStr, cStr, r.op, r.cols, r.key, common.AnyToString(r.t),
	)
}

func (r *row) Driver() Driver {
	return r.d
}

func (r *row) Collection() Collection {
	return r.c
}

func (r *row) Operation() Operation {
	return r.op
}

func (r *row) Columns() Columns {
	return r.cols
}

func (r *row) WithKey(k common.Key) Row {
	r.key = k
	return r
}

func (r *row) WithTimestamp(t time.Time) Row {
	r.t = t
	return r
}

func (r *row) Check() error {
	if r.d == nil {
		return fmt.Errorf("empty driver")
	}
	if r.c == nil {
		return fmt.Errorf("empty collection")
	}
	if r.cols == nil || r.cols.Len() == 0 {
		return fmt.Errorf("empty columns")
	}
	return nil
}
