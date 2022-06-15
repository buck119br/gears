package event

import (
	"fmt"

	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

type Event interface {
	element.Element
	fmt.Stringer
	UnstructuredDataCoder
	UnstructuredDataDecoder

	Meta() Meta
	Data() interface{}

	Local() Local
}
