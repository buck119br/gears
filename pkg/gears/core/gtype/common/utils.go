package common

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

// AnyToString converts value to string
func AnyToString(v interface{}) string {
	switch t := v.(type) {
	case bool:
		return fmt.Sprintf("%t", t)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%f", t)
	case string:
		return t
	case []byte:
		return bytesToString(t)
	case time.Time:
		return v.(time.Time).Format(time.RFC3339)
	default:
		return fmt.Sprintf("%v", t)
	}
}

// AnyToBytes converts value to byte slice
func AnyToBytes(v interface{}) []byte {
	switch t := v.(type) {
	case bool:
		return stringToBytes(fmt.Sprintf("%t", t))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return stringToBytes(fmt.Sprintf("%d", t))
	case float32, float64:
		return stringToBytes(fmt.Sprintf("%f", t))
	case string:
		return stringToBytes(t)
	case []byte:
		return t
	case time.Time:
		return stringToBytes(v.(time.Time).Format(time.RFC3339))
	default:
		return stringToBytes(fmt.Sprintf("%v", t))
	}
}

// bytesToString converts byte slice to string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note: it may break if string and/or slice header will change in the future go versions.
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// stringToBytes converts string to byte slice without memory allocation.
//
// Note: it may break if string and/or slice header will change in the future go versions.
func stringToBytes(s string) []byte {
	strH := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceH := reflect.SliceHeader{
		Data: strH.Data,
		Len:  strH.Len,
		Cap:  strH.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceH))
}
