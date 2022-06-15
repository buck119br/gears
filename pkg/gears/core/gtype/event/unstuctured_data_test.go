package event

import (
	"time"

	. "gopkg.in/check.v1"
)

type StructuredData struct {
	I     int       `sed:"i"`
	I8    int8      `sed:"i8"`
	I16   int16     `sed:"i16"`
	I32   int32     `sed:"i32"`
	I64   int64     `sed:"i64"`
	U     uint      `sed:"u"`
	U8    uint8     `sed:"u8"`
	U16   uint16    `sed:"u16"`
	U32   uint32    `sed:"u32"`
	U64   uint64    `sed:"u64"`
	F32   float32   `sed:"f32"`
	F64   float64   `sed:"f64"`
	Str   string    `sed:"str"`
	Array [3]byte   `sed:"array"`
	Bytes []byte    `sed:"bytes"`
	Time  time.Time `sed:"time"`
}

type UnstructuredDataSuite struct {
	StructuredData    StructuredData
	StructuredDataPtr *StructuredData

	UnstructuredDataForStructuredData    UnstructuredData
	UnstructuredDataForStructuredDataPtr UnstructuredData
}

var _ = Suite(&UnstructuredDataSuite{})

func (us *UnstructuredDataSuite) SetUpSuite(c *C) {
	us.StructuredData = StructuredData{
		I:     -10000,
		I8:    -1,
		I16:   -2,
		I32:   -3,
		I64:   -4,
		U:     10000,
		U8:    1,
		U16:   2,
		U32:   3,
		U64:   4,
		F32:   1,
		F64:   2,
		Str:   "struct_case_string",
		Array: [3]byte{'a', 'b', 'c'},
		Bytes: []byte{'d', 'e', 'f'},
		Time:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
	}

	us.UnstructuredDataForStructuredData = UnstructuredData{
		"i":     "-10000",
		"i8":    "-1",
		"i16":   "-2",
		"i32":   "-3",
		"i64":   "-4",
		"u":     "10000",
		"u8":    "1",
		"u16":   "2",
		"u32":   "3",
		"u64":   "4",
		"f32":   "1",
		"f64":   "2",
		"str":   "struct_case_string",
		"array": "abc",
		"bytes": "def",
		"time":  "2020-01-01T00:00:00+08:00",
	}

	us.StructuredDataPtr = &StructuredData{
		I:     -1000000,
		I8:    -100,
		I16:   -200,
		I32:   -300,
		I64:   -400,
		U:     1000000,
		U8:    100,
		U16:   200,
		U32:   300,
		U64:   400,
		F32:   100,
		F64:   200,
		Str:   "struct_ptr_case_string",
		Array: [3]byte{'u', 'v', 'w'},
		Bytes: []byte{'x', 'y', 'z'},
		Time:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
	}

	us.UnstructuredDataForStructuredDataPtr = UnstructuredData{
		"i":     "-1000000",
		"i8":    "-100",
		"i16":   "-200",
		"i32":   "-300",
		"i64":   "-400",
		"u":     "1000000",
		"u8":    "100",
		"u16":   "200",
		"u32":   "300",
		"u64":   "400",
		"f32":   "100",
		"f64":   "200",
		"str":   "struct_ptr_case_string",
		"array": "uvw",
		"bytes": "xyz",
		"time":  "2021-01-01T00:00:00+08:00",
	}

}

func (us *UnstructuredDataSuite) TestMarshal(c *C) {
	var err error

	r, err := Marshal(us.StructuredData)
	c.Assert(err, IsNil)
	c.Assert(r["i8"], Equals, us.UnstructuredDataForStructuredData["i8"])
	c.Assert(r["i16"], Equals, us.UnstructuredDataForStructuredData["i16"])
	c.Assert(r["i32"], Equals, us.UnstructuredDataForStructuredData["i32"])
	c.Assert(r["i64"], Equals, us.UnstructuredDataForStructuredData["i64"])
	c.Assert(r["u8"], Equals, us.UnstructuredDataForStructuredData["u8"])
	c.Assert(r["u16"], Equals, us.UnstructuredDataForStructuredData["u16"])
	c.Assert(r["u32"], Equals, us.UnstructuredDataForStructuredData["u32"])
	c.Assert(r["u64"], Equals, us.UnstructuredDataForStructuredData["u64"])
	c.Assert(r["f32"], Equals, us.UnstructuredDataForStructuredData["f32"])
	c.Assert(r["f64"], Equals, us.UnstructuredDataForStructuredData["f64"])
	c.Assert(r["str"], Equals, us.UnstructuredDataForStructuredData["str"])
	c.Assert(r["array"], Equals, us.UnstructuredDataForStructuredData["array"])
	c.Assert(r["bytes"], Equals, us.UnstructuredDataForStructuredData["bytes"])
	c.Assert(r["time"], Equals, us.UnstructuredDataForStructuredData["time"])

	r, err = Marshal(us.StructuredDataPtr)
	c.Assert(err, IsNil)
	c.Assert(r["i8"], Equals, us.UnstructuredDataForStructuredDataPtr["i8"])
	c.Assert(r["i16"], Equals, us.UnstructuredDataForStructuredDataPtr["i16"])
	c.Assert(r["i32"], Equals, us.UnstructuredDataForStructuredDataPtr["i32"])
	c.Assert(r["i64"], Equals, us.UnstructuredDataForStructuredDataPtr["i64"])
	c.Assert(r["u8"], Equals, us.UnstructuredDataForStructuredDataPtr["u8"])
	c.Assert(r["u16"], Equals, us.UnstructuredDataForStructuredDataPtr["u16"])
	c.Assert(r["u32"], Equals, us.UnstructuredDataForStructuredDataPtr["u32"])
	c.Assert(r["u64"], Equals, us.UnstructuredDataForStructuredDataPtr["u64"])
	c.Assert(r["f32"], Equals, us.UnstructuredDataForStructuredDataPtr["f32"])
	c.Assert(r["f64"], Equals, us.UnstructuredDataForStructuredDataPtr["f64"])
	c.Assert(r["str"], Equals, us.UnstructuredDataForStructuredDataPtr["str"])
	c.Assert(r["array"], Equals, us.UnstructuredDataForStructuredDataPtr["array"])
	c.Assert(r["bytes"], Equals, us.UnstructuredDataForStructuredDataPtr["bytes"])
	c.Assert(r["time"], Equals, us.UnstructuredDataForStructuredDataPtr["time"])
}

func (us *UnstructuredDataSuite) TestUnmarshal(c *C) {
	var err error

	sc := StructuredData{}
	err = Unmarshal(us.UnstructuredDataForStructuredData, sc)
	c.Assert(err, NotNil)

	spc := &StructuredData{}
	err = Unmarshal(us.UnstructuredDataForStructuredDataPtr, spc)
	c.Assert(err, IsNil)
	c.Assert(spc.I8, Equals, us.StructuredDataPtr.I8)
	c.Assert(spc.I16, Equals, us.StructuredDataPtr.I16)
	c.Assert(spc.I32, Equals, us.StructuredDataPtr.I32)
	c.Assert(spc.I64, Equals, us.StructuredDataPtr.I64)
	c.Assert(spc.U8, Equals, us.StructuredDataPtr.U8)
	c.Assert(spc.U16, Equals, us.StructuredDataPtr.U16)
	c.Assert(spc.U32, Equals, us.StructuredDataPtr.U32)
	c.Assert(spc.U64, Equals, us.StructuredDataPtr.U64)
	c.Assert(spc.F32, Equals, us.StructuredDataPtr.F32)
	c.Assert(spc.F64, Equals, us.StructuredDataPtr.F64)
	c.Assert(spc.Str, Equals, us.StructuredDataPtr.Str)
	c.Assert(spc.Array, Equals, us.StructuredDataPtr.Array)
	c.Assert(spc.Bytes[0], Equals, us.StructuredDataPtr.Bytes[0])
	c.Assert(spc.Bytes[1], Equals, us.StructuredDataPtr.Bytes[1])
	c.Assert(spc.Bytes[2], Equals, us.StructuredDataPtr.Bytes[2])
	c.Assert(spc.Time.Equal(us.StructuredDataPtr.Time), Equals, true)
}
