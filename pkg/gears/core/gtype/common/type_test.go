package common

import (
	"time"

	. "gopkg.in/check.v1"
)

type TypeSuite struct{}

var _ = Suite(&TypeSuite{})

func (ts *TypeSuite) TestType(c *C) {
	b := true
	i := -1
	i8 := int8(-8)
	i16 := int16(-16)
	i32 := int32(-32)
	i64 := int64(-64)
	u := uint(1)
	u8 := uint8(8)
	u16 := uint16(16)
	u32 := uint32(32)
	u64 := uint64(64)
	f32 := float32(1.1111)
	f64 := 2.2222
	str := "abc"
	a := [1]byte{'x'}
	slice := []byte{'y'}
	t := time.Now()

	c.Assert(TypeOf(b), Equals, Bool)
	c.Assert(TypeOf(i), Equals, Integer)
	c.Assert(TypeOf(i8), Equals, Integer)
	c.Assert(TypeOf(i16), Equals, Integer)
	c.Assert(TypeOf(i32), Equals, Integer)
	c.Assert(TypeOf(i64), Equals, Integer)
	c.Assert(TypeOf(u), Equals, Unsigned)
	c.Assert(TypeOf(u8), Equals, Unsigned)
	c.Assert(TypeOf(u16), Equals, Unsigned)
	c.Assert(TypeOf(u32), Equals, Unsigned)
	c.Assert(TypeOf(u64), Equals, Unsigned)
	c.Assert(TypeOf(f32), Equals, Float)
	c.Assert(TypeOf(f64), Equals, Float)
	c.Assert(TypeOf(str), Equals, String)
	c.Assert(TypeOf(a), Equals, Binary)
	c.Assert(TypeOf(slice), Equals, Binary)
	c.Assert(TypeOf(t), Equals, Timestamp)

	c.Assert(TypeOf(&b), Equals, Bool)
	c.Assert(TypeOf(&i), Equals, Integer)
	c.Assert(TypeOf(&i8), Equals, Integer)
	c.Assert(TypeOf(&i16), Equals, Integer)
	c.Assert(TypeOf(&i32), Equals, Integer)
	c.Assert(TypeOf(&i64), Equals, Integer)
	c.Assert(TypeOf(&u), Equals, Unsigned)
	c.Assert(TypeOf(&u8), Equals, Unsigned)
	c.Assert(TypeOf(&u16), Equals, Unsigned)
	c.Assert(TypeOf(&u32), Equals, Unsigned)
	c.Assert(TypeOf(&u64), Equals, Unsigned)
	c.Assert(TypeOf(&f32), Equals, Float)
	c.Assert(TypeOf(&f64), Equals, Float)
	c.Assert(TypeOf(&str), Equals, String)
	c.Assert(TypeOf(&a), Equals, Binary)
	c.Assert(TypeOf(&slice), Equals, Binary)
	c.Assert(TypeOf(&t), Equals, Timestamp)

	c.Assert(TypeOf(any(b)), Equals, Bool)
	c.Assert(TypeOf(any(i)), Equals, Integer)
	c.Assert(TypeOf(any(i8)), Equals, Integer)
	c.Assert(TypeOf(any(i16)), Equals, Integer)
	c.Assert(TypeOf(any(i32)), Equals, Integer)
	c.Assert(TypeOf(any(i64)), Equals, Integer)
	c.Assert(TypeOf(any(u)), Equals, Unsigned)
	c.Assert(TypeOf(any(u8)), Equals, Unsigned)
	c.Assert(TypeOf(any(u16)), Equals, Unsigned)
	c.Assert(TypeOf(any(u32)), Equals, Unsigned)
	c.Assert(TypeOf(any(u64)), Equals, Unsigned)
	c.Assert(TypeOf(any(f32)), Equals, Float)
	c.Assert(TypeOf(any(f64)), Equals, Float)
	c.Assert(TypeOf(any(str)), Equals, String)
	c.Assert(TypeOf(any(a)), Equals, Binary)
	c.Assert(TypeOf(any(slice)), Equals, Binary)
	c.Assert(TypeOf(any(t)), Equals, Timestamp)
}
