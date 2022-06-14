package common

import (
	"time"

	. "gopkg.in/check.v1"
)

type TypeSuite struct{}

var _ = Suite(&TypeSuite{})

func (ts *TypeSuite) TestType(c *C) {
	c.Assert(TypeOf(true), Equals, Bool)
	c.Assert(TypeOf(int(1)), Equals, Integer)
	c.Assert(TypeOf(int8(1)), Equals, Integer)
	c.Assert(TypeOf(int16(1)), Equals, Integer)
	c.Assert(TypeOf(int32(1)), Equals, Integer)
	c.Assert(TypeOf(int64(1)), Equals, Integer)
	c.Assert(TypeOf(uint(1)), Equals, Integer)
	c.Assert(TypeOf(uint8(1)), Equals, Integer)
	c.Assert(TypeOf(uint16(1)), Equals, Integer)
	c.Assert(TypeOf(uint32(1)), Equals, Integer)
	c.Assert(TypeOf(uint64(1)), Equals, Integer)
	c.Assert(TypeOf(float32(1)), Equals, Float)
	c.Assert(TypeOf(float64(1)), Equals, Float)
	c.Assert(TypeOf("1"), Equals, String)
	c.Assert(TypeOf([1]byte{'1'}), Equals, Binary)
	c.Assert(TypeOf([]byte{'1'}), Equals, Binary)
	c.Assert(TypeOf(time.Now()), Equals, Timestamp)

	c.Assert(TypeOf(interface{}(true)), Equals, Bool)
	c.Assert(TypeOf(interface{}(int(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(int8(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(int16(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(int32(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(int64(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(uint(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(uint8(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(uint16(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(uint32(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(uint64(1))), Equals, Integer)
	c.Assert(TypeOf(interface{}(float32(1))), Equals, Float)
	c.Assert(TypeOf(interface{}(float64(1))), Equals, Float)
	c.Assert(TypeOf(interface{}("1")), Equals, String)
	c.Assert(TypeOf(interface{}([1]byte{'1'})), Equals, Binary)
	c.Assert(TypeOf(interface{}([]byte{'1'})), Equals, Binary)
	c.Assert(TypeOf(interface{}(time.Now())), Equals, Timestamp)
}
