package common

import (
	"time"

	. "gopkg.in/check.v1"
)

var _ = Suite(&UtilsSuite{})

type UtilsSuite struct{}

func (us *UtilsSuite) TestAnyToString(c *C) {
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
	f32 := float32(1.1)
	f64 := float64(2.2222)
	str := "abc"
	a := [1]byte{'x'}
	slice := []byte{'y'}
	t := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	c.Assert(AnyToString(b), Equals, "true")
	c.Assert(AnyToString(i), Equals, "-1")
	c.Assert(AnyToString(i8), Equals, "-8")
	c.Assert(AnyToString(i16), Equals, "-16")
	c.Assert(AnyToString(i32), Equals, "-32")
	c.Assert(AnyToString(i64), Equals, "-64")
	c.Assert(AnyToString(u), Equals, "1")
	c.Assert(AnyToString(u8), Equals, "8")
	c.Assert(AnyToString(u16), Equals, "16")
	c.Assert(AnyToString(u32), Equals, "32")
	c.Assert(AnyToString(u64), Equals, "64")
	c.Assert(AnyToString(f32), Equals, "1.1")
	c.Assert(AnyToString(f64), Equals, "2.2222")
	c.Assert(AnyToString(str), Equals, "abc")
	c.Assert(AnyToString(a), Equals, "x")
	c.Assert(AnyToString(slice), Equals, "y")
	c.Assert(AnyToString(t), Equals, "2020-01-01T00:00:00+08:00")

	c.Assert(AnyToString(&b), Equals, "true")
	c.Assert(AnyToString(&i), Equals, "-1")
	c.Assert(AnyToString(&i8), Equals, "-8")
	c.Assert(AnyToString(&i16), Equals, "-16")
	c.Assert(AnyToString(&i32), Equals, "-32")
	c.Assert(AnyToString(&i64), Equals, "-64")
	c.Assert(AnyToString(&u), Equals, "1")
	c.Assert(AnyToString(&u8), Equals, "8")
	c.Assert(AnyToString(&u16), Equals, "16")
	c.Assert(AnyToString(&u32), Equals, "32")
	c.Assert(AnyToString(&u64), Equals, "64")
	c.Assert(AnyToString(&f32), Equals, "1.1")
	c.Assert(AnyToString(&f64), Equals, "2.2222")
	c.Assert(AnyToString(&str), Equals, "abc")
	c.Assert(AnyToString(&a), Equals, "x")
	c.Assert(AnyToString(&slice), Equals, "y")
	c.Assert(AnyToString(&t), Equals, "2020-01-01T00:00:00+08:00")

	c.Assert(AnyToString(interface{}(b)), Equals, "true")
	c.Assert(AnyToString(interface{}(i)), Equals, "-1")
	c.Assert(AnyToString(interface{}(i8)), Equals, "-8")
	c.Assert(AnyToString(interface{}(i16)), Equals, "-16")
	c.Assert(AnyToString(interface{}(i32)), Equals, "-32")
	c.Assert(AnyToString(interface{}(i64)), Equals, "-64")
	c.Assert(AnyToString(interface{}(u)), Equals, "1")
	c.Assert(AnyToString(interface{}(u8)), Equals, "8")
	c.Assert(AnyToString(interface{}(u16)), Equals, "16")
	c.Assert(AnyToString(interface{}(u32)), Equals, "32")
	c.Assert(AnyToString(interface{}(u64)), Equals, "64")
	c.Assert(AnyToString(interface{}(f32)), Equals, "1.1")
	c.Assert(AnyToString(interface{}(f64)), Equals, "2.2222")
	c.Assert(AnyToString(interface{}(str)), Equals, "abc")
	c.Assert(AnyToString(interface{}(a)), Equals, "x")
	c.Assert(AnyToString(interface{}(slice)), Equals, "y")
	c.Assert(AnyToString(interface{}(t)), Equals, "2020-01-01T00:00:00+08:00")
}
