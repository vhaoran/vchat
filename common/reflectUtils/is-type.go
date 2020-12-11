package reflectUtils

import (
	"reflect"
)

func IsPointer(ptr interface{}) bool {
	tp := reflect.TypeOf(ptr)
	switch tp.Kind() {
	case reflect.Ptr, reflect.UnsafePointer:
		return true
	}
	return false
}

//is slice or Point to Slice
func IsSlice(a interface{}) bool {
	tp := reflect.Indirect(reflect.ValueOf(a))
	switch tp.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}

func IsStruct(a interface{}) bool {
	tp := reflect.Indirect(reflect.ValueOf(a))
	switch tp.Kind() {
	case reflect.Struct:
		return true
	}
	return false
}

func IsMap(a interface{}) bool {
	tp := reflect.Indirect(reflect.ValueOf(a))
	switch tp.Kind() {
	case reflect.Map:
		return true
	}
	return false
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	return (rv.Kind() == reflect.Ptr ||
		rv.Kind() == reflect.UnsafePointer) && rv.IsNil()
}

func IsNumber(v interface{}) bool {
	if v == nil {
		return false
	}

	switch v.(type) {
	case int, int8, int16, int32, int64,
		float32, float64,
		uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

func IsInt(v interface{}) bool {
	if v == nil {
		return false
	}

	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return true
	default:
		return false
	}
}

func IsFloat(v interface{}) bool {
	if v == nil {
		return false
	}

	switch v.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}
