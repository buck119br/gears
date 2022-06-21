package graph

import (
	"fmt"
)

// Inbound represents an inbound data link.
type Inbound interface {
	fmt.Stringer

	From() Edge
}

func NewInbound(from Edge) Inbound {
	ib := &inbound{
		from: from,
	}

	return ib
}

type inbound struct {
	from Edge
}

func (ib *inbound) String() string {
	return fmt.Sprintf("from: [%s]", ib.from)
}

func (ib *inbound) From() Edge { return ib.from }

// Outbound represents an outbound data link.
type Outbound interface {
	fmt.Stringer

	To() Edge
}

func NewOutbound(to Edge) Outbound {
	ob := &outbound{
		to: to,
	}

	return ob
}

type outbound struct {
	to Edge
}

func (ob *outbound) String() string {
	return fmt.Sprintf("to: [%s]", ob.to)
}

func (ob *outbound) To() Edge { return ob.to }
