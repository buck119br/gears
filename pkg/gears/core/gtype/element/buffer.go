package element

const (
	DefaultBufferCapacity = 10
)

type Buffer interface {
	Len() int
	Cap() int
	IsEmpty() bool
	IsFull() bool

	// Set sets the Element into Buffer.
	Set(Element)
	// GetAll returns all Elements in Buffer.
	GetAll() []Element
	// Reset resets the Buffer to empty.
	Reset()
}

// NewBuffer returns a Buffer with DefaultBufferCapacity.
// And the underlying buffer will not be allocated until the first Set is called.
// Notice that Buffer is not concurrent-safe.
func NewBuffer(capacity int) Buffer {
	b := new(buffer)
	b.cap = capacity
	b.b = nil
	return b
}

type buffer struct {
	cap int
	b   []Element
}

func (b *buffer) Len() int {
	return len(b.b)
}

func (b *buffer) Cap() int {
	return b.cap
}

func (b *buffer) IsEmpty() bool {
	if len(b.b) == 0 {
		return true
	}
	return false
}

func (b *buffer) IsFull() bool {
	if len(b.b) >= b.cap {
		return true
	}
	return false
}

func (b *buffer) Set(elm Element) {
	if b.b == nil {
		b.allocate()
	}
	if len(b.b) >= b.cap {
		return
	}
	b.b = append(b.b, elm)
}

func (b *buffer) GetAll() []Element {
	elms := make([]Element, len(b.b))
	copy(elms, b.b)
	return elms
}

func (b *buffer) Reset() {
	b.b = nil
}

func (b *buffer) allocate() {
	if b.cap > DefaultBufferCapacity {
		b.b = make([]Element, 0, DefaultBufferCapacity)
	} else {
		b.b = make([]Element, 0, b.cap)
	}
}
