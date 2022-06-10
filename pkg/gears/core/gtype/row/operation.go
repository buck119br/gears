package row

import (
	"fmt"
)

const (
	Default Operation = iota

	Insert
	Query
	Update
	Upsert
	Delete
)

type Operation int

func (o Operation) String() string {
	switch o {
	case Default:
		return "default"

	case Insert:
		return "insert"
	case Query:
		return "query"
	case Update:
		return "update"
	case Upsert:
		return "upsert"
	case Delete:
		return "delete"

	default:
		panic(fmt.Errorf("invalid row operation: [%d]", o))
	}
}
