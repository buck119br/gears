package gears

import (
	"github.com/buck119br/gears/pkg/gears/core/gfunc"
	"github.com/buck119br/gears/pkg/gears/core/graph"
	"github.com/buck119br/gears/pkg/gears/transform/parallel"
)

func Parallel(df Dataflow, o parallel.Option, f DoFunction, inputs []Dataset, args ...any) []Dataset {
	// edge
	fn := gfunc.NewFunction(f)
	n := df.Graph().AddNode(graph.Parallel, fn, args...).WithConcurrency(o.Concurrency())
	// inbounds
	for _, input := range inputs {
		ib := graph.NewInbound(input.Edge())
		n.AddIn(ib)
	}
	// outbounds
	outputDs := make([]Dataset, 0, o.NumOutputs())
	for i := 0; i < o.NumOutputs(); i++ {
		output := df.Graph().AddEdge()

		ob := graph.NewOutbound(output)
		n.AddOut(ob)

		outputD := NewDataset(output)
		outputDs = append(outputDs, outputD)
	}

	return outputDs
}
