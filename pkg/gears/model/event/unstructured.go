package event

import (
	"time"

	"github.com/hamba/avro"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
	"github.com/buck119br/gears/pkg/gears/core/gtype/element"
	"github.com/buck119br/gears/pkg/gears/core/gtype/event"
)

func NewUnstructuredDataEvent() event.Event {
	e := &UnstructuredDataEvent{
		M: event.NewMeta(),
	}
	return e
}

type UnstructuredDataEvent struct {
	M event.Meta             `json:"meta" avro:"meta"`
	D event.UnstructuredData `json:"data" avro:"data"`
	L event.Local            `json:"local"`
}

func (e *UnstructuredDataEvent) Type() element.Type {
	return element.Event
}

func (e *UnstructuredDataEvent) Key() common.Key {
	return e.L.Key()
}

func (e *UnstructuredDataEvent) Timestamp() time.Time {
	return e.M.EventTime()
}

func (e *UnstructuredDataEvent) Code() (event.UnstructuredData, error) {
	return e.D, nil
}

func (e *UnstructuredDataEvent) Decode(d event.UnstructuredData) error {
	e.D = d
	return nil
}

func (e *UnstructuredDataEvent) Meta() event.Meta {
	return e.M
}

func (e *UnstructuredDataEvent) Data() interface{} {
	return e.D
}

func (e *UnstructuredDataEvent) Local() event.Local {
	return e.L
}

var (
	UnstructuredDataEventAvroSchema avro.Schema
)

const (
	UnstructuredDataEventAvroSchemaSpecification = `
	{
		"namespace": "com.github.buck119br.gears.model.event",
		"name": "UnstructuredDataEvent",
		"type": "record",
		"doc": "event with unstructured data",
		"fields": [
			{
				"name": "meta",
				"type": {
					"name": "Meta",
					"type": "record",
					"default": {},
					"doc": "meta object",
					"fields": [
						{
							"name": "event_time",
							"type": {
								"name": "time.Time",
								"type": "long",
								"logicalType": "timestamp-micros",
								"default": 0
							},
							"doc": "time when the event happened, required"
						},
						{
							"name": "id",
							"type": "string",
							"default": "",
							"doc": "event id, optional"
						},
						{
							"name": "event_type",
							"type": "string",
							"default": "",
							"doc": "event type, required"
						},
						{
							"name": "sender",
							"type": "string",
							"default": "",
							"doc": "the one who sent the event, optional"
						}
					]
				}
			},
			{
				"name": "data",
				"type": {
					"name": "UnstructuredData",
					"type": "map",
					"values": "string",
					"default": {}
				},
				"doc": "unstructured data"
			}
		]
	}`
)
