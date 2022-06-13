package event

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/buck119br/gears/pkg/gears/core/gtype/common"
)

const (
	StructuredEventDataTag string = "sed"

	OmitTag string = "-"
)

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

func Marshal(d interface{}) (UnstructuredData, error) {
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

		var fvs string
		switch f.Type {
		default:
			fvs = fmt.Sprintf("%v", fv)
		}

		ud[ft] = fvs
	}

	return ud, nil
}

func Unmarshal(ud UnstructuredData, d interface{}) error {
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
		f := t.Field(i)

		ft := f.Tag.Get(StructuredEventDataTag)
		if len(ft) == 0 || ft == OmitTag {
			continue
		}

		fv := v.Field(i)

		fvs, ok := ud[ft]
		if !ok {
			continue
		}

		switch f.Type.Kind() {
		case reflect.Bool:
			fvr, err := strconv.ParseBool(fvs)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(fvr))

		case reflect.Int8:
			fvr, err := strconv.ParseInt(fvs, 10, 8)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(int8(fvr)))
		case reflect.Int16:
			fvr, err := strconv.ParseInt(fvs, 10, 16)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(int16(fvr)))
		case reflect.Int32:
			fvr, err := strconv.ParseInt(fvs, 10, 32)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(int32(fvr)))
		case reflect.Int64:
			fvr, err := strconv.ParseInt(fvs, 10, 64)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(fvr))

		case reflect.Uint8:
			fvr, err := strconv.ParseUint(fvs, 10, 8)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(uint8(fvr)))
		case reflect.Uint16:
			fvr, err := strconv.ParseUint(fvs, 10, 16)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(uint16(fvr)))
		case reflect.Uint32:
			fvr, err := strconv.ParseUint(fvs, 10, 32)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(uint32(fvr)))
		case reflect.Uint64:
			fvr, err := strconv.ParseUint(fvs, 10, 64)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(uint64(fvr)))

		case reflect.Float32:
			fvr, err := strconv.ParseFloat(fvs, 32)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(float32(fvr)))

		case reflect.Float64:
			fvr, err := strconv.ParseFloat(fvs, 64)
			if err != nil {
				return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
			}
			fv.Set(reflect.ValueOf(fvr))

		case reflect.String:
			fv.Set(reflect.ValueOf(fvs))

		case reflect.Array:
			switch f.Type.Elem().Kind() {
			case reflect.Uint8:
				fv.Set(reflect.ValueOf(common.AnyToBytes(fvs)))
			default:
				return fmt.Errorf("field: [%s] type: [%v] not support", f.Name, f.Type)
			}

		case reflect.Struct:
			switch f.Type.String() {
			case "time.Time":
				fvr, err := time.Parse(fvs, time.RFC3339)
				if err != nil {
					return fmt.Errorf("field: [%s] parse error: [%v] with data: [%s]", f.Name, err, fvs)
				}
				fv.Set(reflect.ValueOf(fvr))
			default:
				return fmt.Errorf("field: [%s] type: [%v] not support", f.Name, f.Type)
			}

		default:
			return fmt.Errorf("field: [%s] type: [%v] not support", f.Name, f.Type)
		}
	}
	return nil
}
