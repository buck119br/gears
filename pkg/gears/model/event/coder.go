package event

import (
	"fmt"

	"github.com/hamba/avro"

	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
	"github.com/buck119br/gears/pkg/gears/core/gtype/event"
)

func UnstructuredDataEventAvroCoder(elm element.Element) ([]byte, error) {
	ude, ok := elm.(*UnstructuredDataEvent)
	if !ok {
		e, ok := elm.(event.Event)
		if !ok {
			return nil, fmt.Errorf("type assertion failure, event expected")
		}
		ud, err := event.Marshal(e.Data())
		if err != nil {
			return nil, fmt.Errorf("marshal data error: [%v]", err)
		}
		ude = &UnstructuredDataEvent{M: e.Meta(), D: ud}
	}

	return avro.Marshal(UnstructuredDataEventAvroSchema, ude)
}
