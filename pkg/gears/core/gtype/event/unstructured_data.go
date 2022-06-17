package event

import (
	"fmt"
	"reflect"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

const (
	StructuredEventDataTag string = "sed"

	OmitTag string = "-"
)

type UnstructuredDataCoder interface {
	Code() (UnstructuredData, error)
}

type UnstructuredDataDecoder interface {
	Decode(UnstructuredData) error
}

type UnstructuredData map[string]string

func (ud UnstructuredData) Union(d UnstructuredData) UnstructuredData {
	union := make(map[string]string)

	for k, v := range ud {
		union[k] = v
	}
	for k, v := range d {
		union[k] = v
	}
	return union
}

func Marshal(d any) (UnstructuredData, error) {
	if d == nil {
		return nil, fmt.Errorf("nil input")
	}

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)

	switch t.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		t = t.Elem()
		v = v.Elem()
	default:
		return nil, fmt.Errorf("invalid data type: [%v], please provide struct or pointer of struct", t)
	}

	return marshalStruct(t, v)
}

func marshalStruct(t reflect.Type, v reflect.Value) (UnstructuredData, error) {
	ud := make(UnstructuredData, 0)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		ft := f.Tag.Get(StructuredEventDataTag)
		if len(ft) == 0 || ft == OmitTag {
			continue
		}

		fv := v.Field(i)
		fvs := common.ReflectValueToString(fv)
		ud[ft] = fvs
	}

	return ud, nil
}

func Unmarshal(ud UnstructuredData, d any) error {
	if ud == nil || len(ud) == 0 {
		return fmt.Errorf("empty raw data")
	}

	if d == nil {
		return fmt.Errorf("nil destination")
	}

	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)

	switch t.Kind() {
	case reflect.Ptr:
		t = t.Elem()
		v = v.Elem()

		return unmarshalStruct(ud, t, v)

	default:
		return fmt.Errorf("invalid destination type: [%v], please provide pointer of struct", t)
	}
}

func unmarshalStruct(ud UnstructuredData, t reflect.Type, v reflect.Value) error {
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)

		fTag := ft.Tag.Get(StructuredEventDataTag)
		if len(fTag) == 0 || fTag == OmitTag {
			continue
		}

		fvs, ok := ud[fTag]
		if !ok {
			continue
		}
		rv, err := common.ParseReflectValue(fvs, ft.Type)
		if err != nil {
			return fmt.Errorf("parse field: [%s] error: [%v]", ft.Name, err)
		}
		fv := v.Field(i)
		fv.Set(rv)
	}
	return nil
}
