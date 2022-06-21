package graph

import (
	"fmt"
	"strings"

	"github.com/buck119br/gears/pkg/gears/core/gfunc"
	"github.com/buck119br/gears/pkg/gears/core/plan"
)

// Graph represents an in-progress deferred execution graph and is easily translatable to the model graph.
// This graph representation allows precise control over scope and connectivity.
type Graph interface {
	fmt.Stringer

	Nodes() []Node
	Edges() []Edge

	AddNode(op Operation, fn gfunc.Function, args ...interface{}) Node
	AddEdge() Edge

	Check() error
	Build() (plan.Plan, error)
}

func NewGraph() Graph {
	g := &graph{
		nodes: make([]Node, 0),
		edges: make([]Edge, 0),
	}

	return g
}

type graph struct {
	nodes []Node
	edges []Edge
}

func (g *graph) String() string {
	nStrings := make([]string, 0, len(g.nodes))
	for _, n := range g.nodes {
		nStrings = append(nStrings, fmt.Sprintf("[%s]", n))
	}
	eStrings := make([]string, 0, len(g.edges))
	for _, e := range g.edges {
		eStrings = append(eStrings, fmt.Sprintf("[%s]", e))
	}

	return fmt.Sprintf(
		"\n\t"+
			"nodes: [\n\t\t%s\n\t]\n\t"+
			"edges: [\n\t\t%s\n\t]\n",
		strings.Join(nStrings, "\n\t\t"), strings.Join(eStrings, "\n\t\t"),
	)
}

func (g *graph) Nodes() []Node {
	return g.nodes
}

func (g *graph) Edges() []Edge {
	return g.edges
}

func (g *graph) AddNode(op Operation, fn gfunc.Function, args ...interface{}) Node {
	n := NewNode(op, fn, args...)
	g.nodes = append(g.nodes, n)
	n.WithId(fmt.Sprintf("%d", len(g.nodes)-1))

	return n
}

func (g *graph) AddEdge() Edge {
	e := NewEdge()
	g.edges = append(g.edges, e)
	e.WithId(fmt.Sprintf("%d", len(g.edges)-1))

	return e
}

func (g *graph) Check() error {
	edges := make(map[Edge]bool)
	for _, e := range g.edges {
		if err := e.Check(); err != nil {
			return fmt.Errorf("edge: [%s] check error: [%v]", e, err)
		}
		edges[e] = true
	}

	reachable := make(map[Edge]Node)
	for _, n := range g.nodes {
		if err := n.Check(); err != nil {
			return fmt.Errorf("node: [%s] check error: [%v]", n, err)
		}
		for _, ib := range n.In() {
			reachable[ib.From()] = n
		}
		for _, ob := range n.Out() {
			reachable[ob.To()] = n
		}
	}

	var ok bool
	for e := range edges {
		if _, ok = reachable[e]; !ok {
			return fmt.Errorf("edge: [%s] is not reachable", e)
		}
	}
	for e := range reachable {
		if _, ok = edges[e]; !ok {
			return fmt.Errorf("edge: [%s] does not belong to this graph", e)
		}
	}

	return nil
}

func (g *graph) Build() (plan.Plan, error) {
	for _, n := range g.nodes {
		maxInboundStage := -1
		for _, ib := range n.In() {
			if err := ib.From().Build(); err != nil {
				return nil, fmt.Errorf("node: [%s] inbound: [%s] build error: [%v]", n, ib, err)
			}
			if maxInboundStage < ib.From().Stage() {
				maxInboundStage = ib.From().Stage()
			}
		}

		nodeStage := maxInboundStage + 1
		n.WithStage(nodeStage)

		for _, ob := range n.Out() {
			if err := ob.To().Build(); err != nil {
				return nil, fmt.Errorf("node: [%s] outbound: [%s] build error: [%v]", n, ob, err)
			}
			ob.To().WithStage(nodeStage)
		}
	}

	for _, e := range g.edges {
		e.WithId(fmt.Sprintf("%d.%s", e.Stage(), e.Id()))
	}

	return nil, nil
}
