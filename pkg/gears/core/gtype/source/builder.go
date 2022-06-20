package source

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type Builder interface {
	Build() (Source, error)

	WithOption(Option) Builder
}

func NewBuilder() Builder {
	b := &builder{}
	return b
}

type builder struct {
	o Option
}

func (b *builder) Build() (Source, error) {
	if err := b.o.Check(); err != nil {
		return nil, fmt.Errorf("option check error: [%v]", err)
	}

	s := &source{
		o:        b.o,
		isClosed: false,
	}

	switch s.o.Mode() {
	case Local:
		s.ch = make(chan element.Element, s.o.BufferSize())
		s.chOk = true
	case Remote:
	default:
	}

	return s, nil
}

func (b *builder) WithOption(o Option) Builder {
	b.o = o
	return b
}
