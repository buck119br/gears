package row

import (
	"fmt"
)

type Driver interface {
	fmt.Stringer

	Model() Model
	Protocol() Protocol
	Instance() string
}

func NewDriver(m Model, p Protocol, instance string) Driver {
	d := &driver{
		m:        m,
		p:        p,
		instance: instance,
	}

	return d
}

type driver struct {
	m        Model
	p        Protocol
	instance string
}

func (d *driver) String() string {
	return fmt.Sprintf("model: [%s], protocol: [%s], instance: [%s]", d.m, d.p, d.instance)
}

func (d *driver) Model() Model {
	return d.m
}

func (d *driver) Protocol() Protocol {
	return d.p
}

func (d *driver) Instance() string {
	return d.instance
}
