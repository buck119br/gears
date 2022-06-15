package event

import (
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type Event interface {
	element.Element
	UnstructuredDataCoder
	UnstructuredDataDecoder

	Meta() Meta
	Data() interface{}

	Local() Local
}
