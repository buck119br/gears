package common

import (
	"reflect"
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
	f64 := 2.2222
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

	c.Assert(AnyToString(any(b)), Equals, "true")
	c.Assert(AnyToString(any(i)), Equals, "-1")
	c.Assert(AnyToString(any(i8)), Equals, "-8")
	c.Assert(AnyToString(any(i16)), Equals, "-16")
	c.Assert(AnyToString(any(i32)), Equals, "-32")
	c.Assert(AnyToString(any(i64)), Equals, "-64")
	c.Assert(AnyToString(any(u)), Equals, "1")
	c.Assert(AnyToString(any(u8)), Equals, "8")
	c.Assert(AnyToString(any(u16)), Equals, "16")
	c.Assert(AnyToString(any(u32)), Equals, "32")
	c.Assert(AnyToString(any(u64)), Equals, "64")
	c.Assert(AnyToString(any(f32)), Equals, "1.1")
	c.Assert(AnyToString(any(f64)), Equals, "2.2222")
	c.Assert(AnyToString(any(str)), Equals, "abc")
	c.Assert(AnyToString(any(a)), Equals, "x")
	c.Assert(AnyToString(any(slice)), Equals, "y")
	c.Assert(AnyToString(any(t)), Equals, "2020-01-01T00:00:00+08:00")
}

func (us *UtilsSuite) TestParseReflectValue(c *C) {
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
	f64 := 2.2222
	str := "abc"
	a := [1]byte{'x'}
	slice := []byte{'y'}
	t := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)

	var (
		rt  reflect.Type
		rv  reflect.Value
		err error
	)

	rt = reflect.TypeOf(b)
	rv, err = ParseReflectValue(AnyToString(b), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface(), Equals, b)

	rt = reflect.TypeOf(i)
	rv, err = ParseReflectValue(AnyToString(i), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(int), Equals, i)
	rt = reflect.TypeOf(i8)
	rv, err = ParseReflectValue(AnyToString(i8), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(int8), Equals, i8)
	rt = reflect.TypeOf(i16)
	rv, err = ParseReflectValue(AnyToString(i16), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(int16), Equals, i16)
	rt = reflect.TypeOf(i32)
	rv, err = ParseReflectValue(AnyToString(i32), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(int32), Equals, i32)
	rt = reflect.TypeOf(i64)
	rv, err = ParseReflectValue(AnyToString(i64), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(int64), Equals, i64)

	rt = reflect.TypeOf(u)
	rv, err = ParseReflectValue(AnyToString(u), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(uint), Equals, u)
	rt = reflect.TypeOf(u8)
	rv, err = ParseReflectValue(AnyToString(u8), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(uint8), Equals, u8)
	rt = reflect.TypeOf(u16)
	rv, err = ParseReflectValue(AnyToString(u16), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(uint16), Equals, u16)
	rt = reflect.TypeOf(u32)
	rv, err = ParseReflectValue(AnyToString(u32), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(uint32), Equals, u32)
	rt = reflect.TypeOf(u64)
	rv, err = ParseReflectValue(AnyToString(u64), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(uint64), Equals, u64)

	rt = reflect.TypeOf(f32)
	rv, err = ParseReflectValue(AnyToString(f32), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(float32), Equals, f32)
	rt = reflect.TypeOf(f64)
	rv, err = ParseReflectValue(AnyToString(f64), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(float64), Equals, f64)

	rt = reflect.TypeOf(str)
	rv, err = ParseReflectValue(str, rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(string), Equals, str)

	rt = reflect.TypeOf(a)
	rv, err = ParseReflectValue(AnyToString(a), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().([1]byte), Equals, a)

	rt = reflect.TypeOf(slice)
	rv, err = ParseReflectValue(AnyToString(slice), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().([]byte)[0], Equals, slice[0])

	rt = reflect.TypeOf(t)
	rv, err = ParseReflectValue(AnyToString(t), rt)
	c.Assert(err, Equals, nil)
	c.Assert(rv.Kind(), Equals, rt.Kind())
	c.Assert(rv.Interface().(time.Time).Equal(t), Equals, true)
}
