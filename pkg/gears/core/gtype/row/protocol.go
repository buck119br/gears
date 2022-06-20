package row

import (
	"fmt"
)

type Protocol interface {
	fmt.Stringer

	Name() string
}

func NewProtocol(name string) Protocol {
	p := &protocol{
		name: name,
	}

	return p
}

type protocol struct {
	name string
}

func (p *protocol) String() string {
	return fmt.Sprintf("name: [%v]", p.name)
}

func (p *protocol) Name() string {
	return p.name
}
