package row

import (
	"fmt"
)

var (
	Relational = NewModel("Relational")
	Analytical = NewModel("Analytical")
	WideColumn = NewModel("WideColumn")
	KeyValue   = NewModel("KeyValue")
	Log        = NewModel("Log")
	File       = NewModel("File")
)

type Model interface {
	fmt.Stringer

	Name() string
}

func NewModel(name string) Model {
	m := new(model)
	m.name = name

	return m
}

type model struct {
	name string
}

func (m *model) String() string {
	return fmt.Sprintf("name: [%v]", m.name)
}

func (m *model) Name() string {
	return m.name
}
