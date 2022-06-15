package event

import (
	"time"

	"github.com/hamba/avro"
	. "gopkg.in/check.v1"

	"github.com/buck119br/gears/pkg/gears/core/gtype/event"
)

var _ = Suite(&CodecSuite{})

type CodecSuite struct{}

func (cs *CodecSuite) SetUpSuite(c *C) {
	var err error
	UnstructuredDataEventAvroSchema, err = avro.Parse(UnstructuredDataEventAvroSchemaSpecification)
	if err != nil {
		c.Fatalf("unstructured data event avro schema parse error: [%v]", err)
	}
}

func (cs *CodecSuite) TestUnstructuredDataEventAvroCodec(c *C) {
	ude := &UnstructuredDataEvent{
		M: event.NewMetaBuilder().
			WithEventTime(time.Now()).
			WithEventType("TestEvent").
			WithSender("gears").
			Build(),
		D: event.UnstructuredData{
			"key_1": "a",
		},
	}

	b, err := UnstructuredDataEventAvroCoder(ude)
	c.Assert(err, Equals, nil)
	udeCopyElm, err := UnstructuredDataEventAvroDecoder(b, nil)
	c.Assert(err, Equals, nil)
	udeCopy := udeCopyElm.(event.Event)

	c.Assert(ude.Meta().EventTime().Equal(udeCopy.Meta().EventTime()), Equals, true)
	c.Assert(ude.Meta().Id(), Equals, udeCopy.Meta().Id())
	c.Assert(ude.Meta().EventType(), Equals, udeCopy.Meta().EventType())
	c.Assert(ude.Meta().Sender(), Equals, udeCopy.Meta().Sender())
	c.Assert(ude.Data().(event.UnstructuredData)["key_1"], Equals, udeCopy.Data().(event.UnstructuredData)["key_1"])
}
