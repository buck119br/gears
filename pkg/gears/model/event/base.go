package event

import (
	"fmt"

	"github.com/hamba/avro"
)

func init() {
	var err error

	UnstructuredDataEventAvroSchema, err = avro.Parse(UnstructuredDataEventAvroSchemaSpecification)
	if err != nil {
		panic(fmt.Errorf("unstructured data event avro schema parse error: [%v]", err))
	}
}
