package parallel

// DefaultOption returns an Option, with numOutputs equals 1 and concurrency equals 1.
func DefaultOption() Option {
	return &option{numOutputs: 1, concurrency: 1}
}

type Option interface {
	NumOutputs() int
	Concurrency() int

	WithNumOutputs(int) Option
	WithConcurrency(int) Option
}

func NewOption() Option {
	o := &option{}

	return o
}

type option struct {
	numOutputs  int
	concurrency int
}

func (o *option) NumOutputs() int {
	return o.numOutputs
}

func (o *option) Concurrency() int {
	return o.concurrency
}

func (o *option) WithNumOutputs(numOutputs int) Option {
	o.numOutputs = numOutputs
	return o
}

func (o *option) WithConcurrency(concurrency int) Option {
	o.concurrency = concurrency
	return o
}
