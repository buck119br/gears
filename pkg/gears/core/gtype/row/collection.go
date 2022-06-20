package row

func NewCollection(name string) Collection {
	c := &collection{
		name: name,
	}

	return c
}

// Collection represents a data collection, such as a table.
type Collection interface {
	Name() string
}

type collection struct {
	name string
}

func (c *collection) Name() string {
	return c.name
}
