package common

import (
	"reflect"
)

var (
	ReflectTypeError = reflect.TypeOf((*error)(nil)).Elem()
)
