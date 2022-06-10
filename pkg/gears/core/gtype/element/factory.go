package element

import (
	"reflect"
	"runtime"
	"strings"
)

type Factory func() Element

func (f Factory) Name() string {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	s := strings.Split(fullName, ".")
	return s[len(s)-1]
}

func DefaultFactory() Element { return nil }
