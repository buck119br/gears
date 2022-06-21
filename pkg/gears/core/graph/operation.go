package graph

import (
	"fmt"
)

const (
	Read Operation = iota
	Combine
	Flatten
	GroupBy
	Join
	Map
	Parallel
	Window
	Write
)

type Operation int

func (o Operation) String() string {
	switch o {
	case Read:
		return "Read"
	case Combine:
		return "Combine"
	case Flatten:
		return "Flatten"
	case GroupBy:
		return "GroupBy"
	case Join:
		return "Join"
	case Map:
		return "Map"
	case Parallel:
		return "Parallel"
	case Window:
		return "Window"
	case Write:
		return "Write"
	default:
		panic(fmt.Errorf("invalid operation: [%d]", o))
	}
}

func (o Operation) NodeType() NodeType {
	switch o {
	case Read:
		return Input
	case Combine, Flatten, GroupBy, Join, Map, Parallel, Window:
		return Transform
	case Write:
		return Output
	default:
		panic(fmt.Errorf("invalid operation: [%d]", o))
	}
}

func (o Operation) Check() error {
	if o < 0 || o > Write {
		return fmt.Errorf("invalid operation: [%d]", o)
	}

	return nil
}
