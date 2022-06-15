package row

import (
	"fmt"
	"strings"
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
	Columns() []Column

	WithKey(common.Key) Row
	WithTimestamp(time.Time) Row

	Check() error
}

func NewRow(d Driver, c Collection, op Operation, cols []Column) Row {
	r := newRow()
	r.d = d
	r.c = c
	r.op = op
	r.cols = cols

	return r
}

type row struct {
	d    Driver
	c    Collection
	op   Operation
	cols []Column

	key common.Key
	t   time.Time
}

func newRow() *row {
	r := new(row)

	return r
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

	colsStrs := make([]string, 0, len(r.cols))
	for _, col := range r.cols {
		colsStrs = append(colsStrs, col.String())
	}

	return fmt.Sprintf("driver: [%s], collection: [%s], operation: [%s], columns: [%s], key: [%s], timestamp: [%s]",
		dStr, cStr, r.op, strings.Join(colsStrs, ", "), r.key, common.AnyToString(r.t),
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

func (r *row) Columns() []Column {
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
	if r.Driver() == nil {
		return fmt.Errorf("empty driver")
	}
	if r.Collection() == nil {
		return fmt.Errorf("empty collection")
	}
	if r.Columns() == nil || len(r.Columns()) == 0 {
		return fmt.Errorf("empty columns")
	}
	return nil
}
