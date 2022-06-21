package gears

import (
	"fmt"
)

type Option interface {
	fmt.Stringer

	Name() string

	WithName(name string) Option
}

func NewOption() Option {
	o := &option{}

	return o
}

type option struct {
	name string
}

func (o *option) String() string {
	return fmt.Sprintf("name: [%s]", o.name)
}

func (o *option) Name() string { return o.name }

func (o *option) WithName(name string) Option {
	o.name = name
	return o
}
