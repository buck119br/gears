package graph

import (
	"fmt"
)

type Edge interface {
	fmt.Stringer

	Id() string
	// Bounded returns whether the dateset is bounded.
	Bounded() bool
	Stage() int

	WithId(string) Edge
	WithBounded(bool) Edge
	WithStage(int) Edge

	Check() error
	Build() error
}

func NewEdge() Edge {
	e := &edge{}

	return e
}

type edge struct {
	id      string
	bounded bool
	stage   int
}

func (e *edge) String() string {
	return fmt.Sprintf("id: [%s], bounded: [%t], stage: [%d]", e.id, e.bounded, e.stage)
}

func (e *edge) Id() string {
	return e.id
}

func (e *edge) Bounded() bool {
	return e.bounded
}

func (e *edge) Stage() int {
	return e.stage
}

func (e *edge) WithId(id string) Edge {
	e.id = id
	return e
}

func (e *edge) WithBounded(bounded bool) Edge {
	e.bounded = bounded
	return e
}

func (e *edge) WithStage(stage int) Edge {
	e.stage = stage
	return e
}

func (e *edge) Check() error {
	return nil
}

func (e *edge) Build() error {
	return nil
}
