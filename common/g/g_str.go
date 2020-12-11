package g

import (
	"errors"
	"reflect"
	"strings"
)

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyAll(l ...interface{}) bool {
	for _, v := range l {
		if !IsEmpty(v) {
			return false
		}
	}
	return true
}

func IsEmptyAllInfo(info string, l ...interface{}) error {
	for _, v := range l {
		if !IsEmpty(v) {
			return nil
		}
	}

	return errors.New(info)
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyOr(l ...interface{}) bool {
	for _, v := range l {
		if IsEmpty(v) {
			return true
		}
	}
	return false
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmptyOrMsg(msg string, l ...interface{}) (err error) {
	if IsEmptyOr(l) {
		return errors.New(msg)
	}
	return
}

func IsEmptyOrInfo(info string, l ...interface{}) (err error) {
	all := strings.Split(info, "/")
	s := ""

	for i, v := range l {
		if IsEmpty(v) {
			if i >= 0 && i < len(all) {
				s += "/" + all[i]
			} else {
				if s == "" {
					s += "/不合法数据项"
				}
			}
		}
	}

	if len(s) > 0 {
		return errors.New("数据不合法:" + s)
	}
	return nil
}

//str len is 0/struct is nil/array/chan/slice is nil or len is 0
func IsEmpty(dst interface{}) bool {
	val := reflect.ValueOf(dst)

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			return IsZero(val.Int())
		}
	case reflect.Float32, reflect.Float64:
		{
			return IsZero(val.Float())
		}
	case reflect.String:
		{
			return val.Len() == 0
		}
	case reflect.Map, reflect.Slice, reflect.Array:
		{ //len == 0
			return val.Len() == 0
		}
	case reflect.Chan:
		{ //len == 0
			return val.Len() == 0
		}
	case reflect.Struct:
		{ //len == 0
			if val.IsValid() {
				return true
			}
			return false
		}
	case reflect.Ptr, reflect.Uintptr:
		{
			return val.IsNil()
		}
	}

	return false
}
