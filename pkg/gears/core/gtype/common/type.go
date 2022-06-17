package common

import (
	"fmt"
	"reflect"
	"time"
)

const (
	Bool Type = iota
	Integer
	Unsigned
	Float
	String
	Binary
	Timestamp
)

type Type int

func (t Type) String() string {
	switch t {
	case Bool:
		return "Bool"
	case Integer:
		return "Integer"
	case Unsigned:
		return "Unsigned"
	case Float:
		return "Float"
	case String:
		return "String"
	case Binary:
		return "Binary"
	case Timestamp:
		return "Timestamp"
	default:
		panic(fmt.Errorf("invalid type: [%d]", t))
	}
}

// TypeOf gets the Type of v.
// Only part of Types are supported.
func TypeOf(v any) Type {
	return TypeOfReflectType(reflect.TypeOf(v))
}

// TypeOfReflectType returns the Type of reflect.Type t.
func TypeOfReflectType(t reflect.Type) Type {
	switch t.Kind() {
	case reflect.Bool:
		return Bool

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Integer
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Unsigned

	case reflect.Float32, reflect.Float64:
		return Float

	case reflect.String:
		return String

	case reflect.Array, reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Uint8:
			return Binary
			// TODO: other array or slice type
		default:
			panic(fmt.Errorf("invalid type: [%s]", t))
		}

	case reflect.Struct:
		switch t.String() {
		case "time.Time":
			return Timestamp
			// TODO: other struct type
		default:
			panic(fmt.Errorf("invalid type: [%s]", t))
		}

	case reflect.Pointer, reflect.Interface:
		return TypeOfReflectType(t.Elem())

	default:
		panic(fmt.Errorf("invalid reflect type: [%s]", t))
	}
}

func New(t Type) any {
	switch t {
	case Bool:
		return false
	case Integer:
		return int64(0)
	case Unsigned:
		return uint64(0)
	case Float:
		return float64(0)
	case String:
		return ""
	case Binary:
		return []byte{}
	case Timestamp:
		return time.Time{}
	default:
		panic(fmt.Errorf("invalid type: [%s]", t))
	}
}
