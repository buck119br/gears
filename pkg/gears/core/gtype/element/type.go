package element

import (
	"fmt"
)

const (
	KV Type = iota
	Event
	Row
	WindowedValues
	Watermark
	Signal
)

type Type int

func (t Type) String() string {
	switch t {
	case KV:
		return "KV"
	case Event:
		return "Event"
	case Row:
		return "Row"
	case WindowedValues:
		return "WindowedValues"
	case Watermark:
		return "Watermark"
	case Signal:
		return "Signal"
	default:
		panic(fmt.Errorf("invalid element type: [%d]", t))
	}
}
