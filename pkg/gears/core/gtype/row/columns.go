package row

import (
	"fmt"
	"strings"
)

type Columns interface {
	fmt.Stringer

	Len() int
	Add(Column) Columns
	Get() []Column
}

func NewColumns() Columns {
	c := &columns{}

	return c
}

type columns struct {
	cols []Column
}

func (c *columns) String() string {
	colsStrs := make([]string, 0, len(c.cols))
	for _, col := range c.cols {
		colsStrs = append(colsStrs, fmt.Sprintf("%s", col))
	}
	return strings.Join(colsStrs, ", ")
}

func (c *columns) Len() int {
	return len(c.cols)
}

func (c *columns) Add(col Column) Columns {
	c.cols = append(c.cols, col)
	return c
}

func (c *columns) Get() []Column {
	return c.cols
}
