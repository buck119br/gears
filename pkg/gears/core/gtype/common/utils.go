package common

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

// AnyToString converts value to string
func AnyToString(v interface{}) string {
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

// AnyToBytes converts value to byte slice
func AnyToBytes(v interface{}) []byte {
	return StringToBytes(AnyToString(v))
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
