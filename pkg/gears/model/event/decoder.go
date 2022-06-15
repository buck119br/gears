package event

import (
	"fmt"

	"github.com/hamba/avro"

	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
)

func UnstructuredDataEventAvroDecoder(raw []byte, elm element.Element) (element.Element, error) {
	var e element.Element
	var ok bool
	if elm != nil {
		e, ok = elm.(*UnstructuredDataEvent)
		if !ok {
			return nil, fmt.Errorf("type assertion failure, unstructured data event expected")
		}
	} else {
		e = NewUnstructuredDataEvent()
	}
	if err := avro.Unmarshal(UnstructuredDataEventAvroSchema, raw, e); err != nil {
		return nil, fmt.Errorf("unmarshal error: [%v]", err)
	}
	return e, nil
}
