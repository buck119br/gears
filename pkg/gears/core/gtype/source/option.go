package source

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMode = errors.New("invalid mode")
)

type Option interface {
	fmt.Stringer

	Mode() Mode
	BufferSize() int

	WithMode(Mode) Option

	Check() error
}

func NewOption() Option {
	o := &option{
		m:          Local,
		bufferSize: 0,
	}
	return o
}

type option struct {
	m          Mode
	bufferSize int
}

func (o *option) String() string {
	return fmt.Sprintf("mode: [%s], buffer size: [%d]", o.m, o.bufferSize)
}

func (o *option) Mode() Mode {
	return o.m
}

func (o *option) BufferSize() int {
	return o.bufferSize
}

func (o *option) WithMode(m Mode) Option {
	o.m = m
	return o
}

func (o *option) Check() error {
	switch o.m {
	case Local, Remote:
	default:
		return ErrInvalidMode
	}

	return nil
}

const (
	Local Mode = iota
	Remote
)

type Mode int

func (m Mode) String() string {
	switch m {
	case Local:
		return "Local"
	case Remote:
		return "Remote"
	default:
		panic(fmt.Errorf("invalid source mode: [%d]", m))
	}
}
