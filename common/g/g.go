package g

import (
	"math"
	"reflect"
)

func IsString(i interface{}) bool {
	k := reflect.TypeOf(i).Kind()
	return k == reflect.String
}

const (
	constZERO = 0.00000000001
)

//new
func IsZeroAll(intOrFloatOrPtr ...interface{}) bool {
	for _, v := range intOrFloatOrPtr {
		if !IsZero(v) {
			return false
		}
	}
	return true
}

//new
func IsZeroOr(intOrFloatOrPtr ...interface{}) bool {
	for _, v := range intOrFloatOrPtr {
		if IsZero(v) {
			return true
		}
	}
	return false
}

//new
func IsZero(intOrFloatOrPtr interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(intOrFloatOrPtr))

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			return val.Int() == 0
		}
	case reflect.Float32, reflect.Float64:
		{
			f := val.Float()
			return math.Abs(f) < constZERO
		}
	}
	return false
}
