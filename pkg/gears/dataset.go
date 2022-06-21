package gears

import (
	"github.com/buck119br/gears/pkg/gears/core/graph"
)

// Dataset is a collection of Element.
// A Dataset can contain either a bounded or unbounded number of elements.
// Bounded and unbounded Dataset are produced as the output of Transform, and can be passed as the inputs of other Transform.
// Some root Transforms produce bounded Dataset and others produce unbounded ones.
//
// Each element in a Dataset has an associated timestamp.
// Sources assign timestamps to elements when they create Dataset,
// and other Transforms propagate these timestamps from their input to their output implicitly or explicitly.
type Dataset interface {
	Edge() graph.Edge
}

func NewDataset(e graph.Edge) Dataset {
	ds := &dataset{
		e: e,
	}

	return ds
}

type dataset struct {
	e graph.Edge
}

func (ds *dataset) Edge() graph.Edge {
	return ds.e
}
