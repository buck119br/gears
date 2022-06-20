package gfunc

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Function interface {
	Name() string
	NumIn() int
	NumOut() int

	Call(args ...any) ([]reflect.Value, error)
}

func NewFunction(fn any) Function {
	f := &function{
		fn: reflect.ValueOf(fn),
	}

	if f.fn.Type().Kind() != reflect.Func {
		panic(fmt.Errorf("function expected"))
	}

	return f
}

type function struct {
	fn reflect.Value
}

func (f *function) Name() string {
	fullName := runtime.FuncForPC(f.fn.Pointer()).Name()
	s := strings.Split(fullName, "/")
	return s[len(s)-1]
}

func (f *function) NumIn() int {
	return f.fn.Type().NumIn()
}

func (f *function) NumOut() int {
	return f.fn.Type().NumOut()
}

func (f *function) Call(args ...any) ([]reflect.Value, error) {
	argValues := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		argValues = append(argValues, reflect.ValueOf(arg))
	}

	return f.fn.Call(argValues), nil
}
