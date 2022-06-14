package event

import (
	"time"

	. "gopkg.in/check.v1"
)

type CodecTestStructCase struct {
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

type CodecSuite struct {
	StructCase                            CodecTestStructCase
	StructPtrCase                         *CodecTestStructCase
	UnstructuredEventDataForStructCase    UnstructuredData
	UnstructuredEventDataForStructPtrCase UnstructuredData
}

var _ = Suite(&CodecSuite{})

func (cs *CodecSuite) SetUpSuite(c *C) {
	cs.StructCase = CodecTestStructCase{
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

	cs.UnstructuredEventDataForStructCase = UnstructuredData{
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

	cs.StructPtrCase = &CodecTestStructCase{
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

	cs.UnstructuredEventDataForStructPtrCase = UnstructuredData{
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

func (cs *CodecSuite) TestCodec(c *C) {
	var err error

	r, err := Marshal(cs.StructCase)
	c.Assert(err, IsNil)
	c.Assert(r["i8"], Equals, cs.UnstructuredEventDataForStructCase["i8"])
	c.Assert(r["i16"], Equals, cs.UnstructuredEventDataForStructCase["i16"])
	c.Assert(r["i32"], Equals, cs.UnstructuredEventDataForStructCase["i32"])
	c.Assert(r["i64"], Equals, cs.UnstructuredEventDataForStructCase["i64"])
	c.Assert(r["u8"], Equals, cs.UnstructuredEventDataForStructCase["u8"])
	c.Assert(r["u16"], Equals, cs.UnstructuredEventDataForStructCase["u16"])
	c.Assert(r["u32"], Equals, cs.UnstructuredEventDataForStructCase["u32"])
	c.Assert(r["u64"], Equals, cs.UnstructuredEventDataForStructCase["u64"])
	c.Assert(r["f32"], Equals, cs.UnstructuredEventDataForStructCase["f32"])
	c.Assert(r["f64"], Equals, cs.UnstructuredEventDataForStructCase["f64"])
	c.Assert(r["str"], Equals, cs.UnstructuredEventDataForStructCase["str"])
	c.Assert(r["array"], Equals, cs.UnstructuredEventDataForStructCase["array"])
	c.Assert(r["bytes"], Equals, cs.UnstructuredEventDataForStructCase["bytes"])
	c.Assert(r["time"], Equals, cs.UnstructuredEventDataForStructCase["time"])

	r, err = Marshal(cs.StructPtrCase)
	c.Assert(err, IsNil)
	c.Assert(r["i8"], Equals, cs.UnstructuredEventDataForStructPtrCase["i8"])
	c.Assert(r["i16"], Equals, cs.UnstructuredEventDataForStructPtrCase["i16"])
	c.Assert(r["i32"], Equals, cs.UnstructuredEventDataForStructPtrCase["i32"])
	c.Assert(r["i64"], Equals, cs.UnstructuredEventDataForStructPtrCase["i64"])
	c.Assert(r["u8"], Equals, cs.UnstructuredEventDataForStructPtrCase["u8"])
	c.Assert(r["u16"], Equals, cs.UnstructuredEventDataForStructPtrCase["u16"])
	c.Assert(r["u32"], Equals, cs.UnstructuredEventDataForStructPtrCase["u32"])
	c.Assert(r["u64"], Equals, cs.UnstructuredEventDataForStructPtrCase["u64"])
	c.Assert(r["f32"], Equals, cs.UnstructuredEventDataForStructPtrCase["f32"])
	c.Assert(r["f64"], Equals, cs.UnstructuredEventDataForStructPtrCase["f64"])
	c.Assert(r["str"], Equals, cs.UnstructuredEventDataForStructPtrCase["str"])
	c.Assert(r["array"], Equals, cs.UnstructuredEventDataForStructPtrCase["array"])
	c.Assert(r["bytes"], Equals, cs.UnstructuredEventDataForStructPtrCase["bytes"])
	c.Assert(r["time"], Equals, cs.UnstructuredEventDataForStructPtrCase["time"])

	sc := CodecTestStructCase{}
	err = Unmarshal(cs.UnstructuredEventDataForStructCase, sc)
	c.Assert(err, NotNil)

	spc := &CodecTestStructCase{}
	err = Unmarshal(cs.UnstructuredEventDataForStructPtrCase, spc)
	c.Assert(err, IsNil)
	c.Assert(spc.I8, Equals, cs.StructPtrCase.I8)
	c.Assert(spc.I16, Equals, cs.StructPtrCase.I16)
	c.Assert(spc.I32, Equals, cs.StructPtrCase.I32)
	c.Assert(spc.I64, Equals, cs.StructPtrCase.I64)
	c.Assert(spc.U8, Equals, cs.StructPtrCase.U8)
	c.Assert(spc.U16, Equals, cs.StructPtrCase.U16)
	c.Assert(spc.U32, Equals, cs.StructPtrCase.U32)
	c.Assert(spc.U64, Equals, cs.StructPtrCase.U64)
	c.Assert(spc.F32, Equals, cs.StructPtrCase.F32)
	c.Assert(spc.F64, Equals, cs.StructPtrCase.F64)
	c.Assert(spc.Str, Equals, cs.StructPtrCase.Str)
	c.Assert(spc.Array, Equals, cs.StructPtrCase.Array)
	c.Assert(spc.Bytes[0], Equals, cs.StructPtrCase.Bytes[0])
	c.Assert(spc.Bytes[1], Equals, cs.StructPtrCase.Bytes[1])
	c.Assert(spc.Bytes[2], Equals, cs.StructPtrCase.Bytes[2])
	c.Assert(spc.Time.Equal(cs.StructPtrCase.Time), Equals, true)
}
