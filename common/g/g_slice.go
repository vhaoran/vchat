package g

import (
	"reflect"
)

func InSlice(src interface{}, l interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(l))
	switch v.Type().Kind() {
	case reflect.Slice, reflect.Array:
		{
			for i := 0; i < v.Len(); i++ {
				elem := reflect.ValueOf(v.Index(i).Interface())
				//if src == reflect.Indirect(elem).Interface() {
				if reflect.DeepEqual(src, reflect.Indirect(elem).Interface()) {
					return true
				}
			}
		}
	default:
		//do nothing
	}
	return false
}

func In(src interface{}, l ...interface{}) bool {
	for _, v := range l {
		if src == v {
			return true
		}
	}
	return false
}
