package source

import (
	"errors"
	"runtime"
	"sync"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

var (
	ErrClosed     = errors.New("source is closed")
	ErrNilElement = errors.New("nil element")
)

type Source interface {
	common.Closer

	Emit(element.Element) error

	HasNext() bool
	Next()
	Current() element.Element
}

type source struct {
	mu sync.RWMutex

	o Option

	isClosed bool

	ch   chan element.Element
	chOk bool

	elm element.Element
}

func (s *source) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isClosed {
		return ErrClosed
	}

	s.isClosed = true
	switch s.o.Mode() {
	case Local:
		for len(s.ch) != 0 {
			s.mu.Unlock()
			runtime.Gosched()
			s.mu.Lock()
		}
		close(s.ch)
		return nil
	case Remote:
		return nil
	default:
		return nil
	}
}

func (s *source) Emit(elm element.Element) error {
	if elm == nil {
		return ErrNilElement
	}

	s.mu.Lock()

	if s.isClosed {
		s.mu.Unlock()
		return ErrClosed
	}

	s.mu.Unlock()
	switch s.o.Mode() {
	case Local:
		s.ch <- elm
		return nil
	case Remote:
		return nil
	default:
		return nil
	}
}

func (s *source) HasNext() bool {
	switch s.o.Mode() {
	case Local:
		s.elm, s.chOk = <-s.ch
		return s.chOk
	case Remote:
		return false
	default:
		return false
	}
}

func (s *source) Next() {
	switch s.o.Mode() {
	case Local:
	case Remote:
	default:
	}
}

func (s *source) Current() element.Element {
	s.mu.RLock()
	defer s.mu.RUnlock()

	switch s.o.Mode() {
	case Local:
		return s.elm
	case Remote:
		return nil
	default:
		return nil
	}
}
