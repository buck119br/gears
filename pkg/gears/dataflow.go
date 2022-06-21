package gears

import (
	"context"
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/graph"
	"github.com/buck119br/gears/pkg/gears/log"
)

// Dataflow manages a DAG (Directed Acyclic Graph) of primitive Transforms, and the Dataset that the Transforms consume and produce.
// Each Dataflow is self-contained and isolated from any other Dataflow.
// The Dataflow owns the Datasets and Transforms, and they can be used by that Dataflow only.
type Dataflow interface {
	fmt.Stringer

	Graph() graph.Graph
	Run(context.Context) error
}

func NewDataflow(o Option) Dataflow {
	df := &dataflow{
		o: o,
		g: graph.NewGraph(),
	}

	return df
}

type dataflow struct {
	o Option
	g graph.Graph
}

func (df *dataflow) String() string {
	return fmt.Sprintf(
		"\n"+
			"option: [%s]\n"+
			"graph: [%s]\n",
		df.o, df.g,
	)
}

func (df *dataflow) Graph() graph.Graph { return df.g }

func (df *dataflow) Run(ctx context.Context) error {
	var err error

	if err = df.g.Check(); err != nil {
		return fmt.Errorf("graph check error: [%v]", err)
	}

	p, err := df.g.Build()
	if err != nil {
		return fmt.Errorf("graph build error: [%v]", err)
	}

	log.Infof("dataflow: [%s] build finished", df)

	if err = p.Run(); err != nil {
		return fmt.Errorf("plan run error: [%v]", err)
	}

	return nil
}
