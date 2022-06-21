package graph

import (
	"fmt"
	"strings"

	"github.com/buck119br/gears/pkg/gears/core/gfunc"
)

type Node interface {
	fmt.Stringer

	Id() string
	Type() NodeType
	Stage() int

	Operation() Operation
	Function() gfunc.Function
	Arguments() []any
	In() []Inbound
	Out() []Outbound
	Concurrency() int

	WithId(string) Node
	WithStage(int) Node

	AddIn(...Inbound) Node
	AddOut(...Outbound) Node
	WithConcurrency(int) Node

	Check() error
	Build() error
}

func NewNode(op Operation, fn gfunc.Function, args ...any) Node {
	n := &node{
		t:     op.NodeType(),
		stage: int(op.NodeType()),
		op:    op,
		fn:    fn,
		args:  args,
	}
	return n
}

type node struct {
	id    string
	t     NodeType
	stage int

	op          Operation
	fn          gfunc.Function
	args        []any
	in          []Inbound
	out         []Outbound
	concurrency int
}

func (n *node) String() string {
	ibIds := make([]string, 0, len(n.in))
	for _, ib := range n.in {
		ibIds = append(ibIds, ib.From().Id())
	}
	obIds := make([]string, 0, len(n.out))
	for _, ob := range n.out {
		obIds = append(obIds, ob.To().Id())
	}
	fnStr := fmt.Sprintf("id: [%s], type: [%s], stage: [%d], operation: [%s], function: [%s], concurrency: [%d]",
		n.id, n.t, n.stage, n.op, n.fn.Name(), n.concurrency)

	return fmt.Sprintf("inbounds: [%s] -> node: [%s] -> outbounds: [%s]", strings.Join(ibIds, ", "), fnStr, strings.Join(obIds, ", "))
}

func (n *node) Id() string {
	return n.id
}

func (n *node) Type() NodeType {
	return n.t
}

func (n *node) Stage() int {
	return n.stage
}

func (n *node) Operation() Operation {
	return n.op
}

func (n *node) Function() gfunc.Function {
	return n.fn
}

func (n *node) Arguments() []any {
	return n.args
}

func (n *node) In() []Inbound {
	return n.in
}

func (n *node) Out() []Outbound {
	return n.out
}

func (n *node) Concurrency() int {
	return n.concurrency
}

func (n *node) WithId(id string) Node {
	n.id = id
	return n
}

func (n *node) WithStage(s int) Node {
	n.stage = s
	return n
}

func (n *node) AddIn(in ...Inbound) Node {
	n.in = append(n.in, in...)
	return n
}

func (n *node) AddOut(out ...Outbound) Node {
	n.out = append(n.out, out...)
	return n
}

func (n *node) WithConcurrency(c int) Node {
	n.concurrency = c
	return n
}

func (n *node) Check() error {
	err := n.op.Check()
	if err != nil {
		return fmt.Errorf("operation check error: [%v]", err)
	}
	if n.fn == nil {
		return fmt.Errorf("empty function")
	}

	return nil
}

func (n *node) Build() error {
	return nil
}

const (
	Input NodeType = iota
	Transform
	Output
)

type NodeType int

func (t NodeType) String() string {
	switch t {
	case Input:
		return "Input"
	case Transform:
		return "Transform"
	case Output:
		return "Output"
	default:
		panic(fmt.Errorf("invalid node type: [%d]", t))
	}
}

func (t NodeType) Check() error {
	if t < 0 || t > Output {
		return fmt.Errorf("invalid node type: [%d]", t)
	}

	return nil
}
