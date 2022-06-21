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
	e := &multiEdge{}

	return e
}

type multiEdge struct {
	id      string
	bounded bool
	stage   int
}

func (e *multiEdge) String() string {
	return fmt.Sprintf("id: [%s], bounded: [%t], stage: [%d]", e.id, e.bounded, e.stage)
}

func (e *multiEdge) Id() string {
	return e.id
}

func (e *multiEdge) Bounded() bool {
	return e.bounded
}

func (e *multiEdge) Stage() int {
	return e.stage
}

func (e *multiEdge) WithId(id string) Edge {
	e.id = id
	return e
}

func (e *multiEdge) WithBounded(bounded bool) Edge {
	e.bounded = bounded
	return e
}

func (e *multiEdge) WithStage(stage int) Edge {
	e.stage = stage
	return e
}

func (e *multiEdge) Check() error {
	return nil
}

func (e *multiEdge) Build() error {
	return nil
}
