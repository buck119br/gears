package common

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

// AnyToBytes converts value to byte slice
func AnyToBytes(v any) []byte {
	return StringToBytes(AnyToString(v))
}

// AnyToString converts value to string
func AnyToString(v any) string {
	return ReflectValueToString(reflect.ValueOf(v))
}

func ReflectValueToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())

	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)

	case reflect.String:
		return v.String()

	case reflect.Array, reflect.Slice:
		switch v.Type().Elem().Kind() {
		case reflect.Uint8:
			uint8s := make([]uint8, v.Cap())
			for i := 0; i < v.Cap(); i++ {
				uint8s[i] = uint8(v.Index(i).Uint())
			}
			return BytesToString(uint8s)
			// TODO: other array or slice type
		default:
			return fmt.Sprintf("%v", v)
		}

	case reflect.Struct:
		switch v.Type().String() {
		case "time.Time":
			return v.Interface().(time.Time).Format(time.RFC3339Nano)
			// TODO: other struct type
		default:
			return fmt.Sprintf("%v", v)
		}

	case reflect.Pointer, reflect.Interface:
		return ReflectValueToString(v.Elem())

	default:
		return fmt.Sprintf("%v", v)
	}
}

// BytesToString converts byte slice to string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note: it may break if string and/or slice header will change in the future go versions.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice without memory allocation.
//
// Note: it may break if string and/or slice header will change in the future go versions.
func StringToBytes(s string) []byte {
	strH := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceH := reflect.SliceHeader{
		Data: strH.Data,
		Len:  strH.Len,
		Cap:  strH.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceH))
}

func ParseReflectValue(s string, t reflect.Type) (reflect.Value, error) {
	var fv reflect.Value
	switch t.Kind() {
	case reflect.Bool:
		v, err := strconv.ParseBool(s)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(v), nil

	case reflect.Int:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(int(v)), nil
	case reflect.Int8:
		v, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(int8(v)), nil
	case reflect.Int16:
		v, err := strconv.ParseInt(s, 10, 16)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(int16(v)), nil
	case reflect.Int32:
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(int32(v)), nil
	case reflect.Int64:
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(v), nil

	case reflect.Uint:
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(uint(v)), nil
	case reflect.Uint8:
		v, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(uint8(v)), nil
	case reflect.Uint16:
		v, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(uint16(v)), nil
	case reflect.Uint32:
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(uint32(v)), nil
	case reflect.Uint64:
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(uint64(v)), nil

	case reflect.Float32:
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(float32(v)), nil
	case reflect.Float64:
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
		}
		return reflect.ValueOf(v), nil

	case reflect.String:
		return reflect.ValueOf(s), nil

	case reflect.Array:
		switch t.Elem().Kind() {
		case reflect.Uint8:
			arrayType := reflect.ArrayOf(len(s), t.Elem())
			array := reflect.New(arrayType).Elem()
			for i, ch := range s {
				array.Index(i).Set(reflect.ValueOf(byte(ch)))
			}
			return array, nil
		default:
			return fv, fmt.Errorf("type: [%v] not support", t)
		}

	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Uint8:
			return reflect.ValueOf(StringToBytes(s)), nil
		default:
			return fv, fmt.Errorf("type: [%v] not support", t)
		}

	case reflect.Struct:
		switch t.String() {
		case "time.Time":
			v, err := time.Parse(time.RFC3339Nano, s)
			if err != nil {
				return fv, fmt.Errorf("parse error: [%v] with data: [%s]", err, s)
			}
			return reflect.ValueOf(v), nil
		default:
			return fv, fmt.Errorf("type: [%v] not support", t)
		}

	default:
		return fv, fmt.Errorf("type: [%v] not support", t)
	}
}
