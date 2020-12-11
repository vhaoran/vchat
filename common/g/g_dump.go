package g

import (
	"fmt"
	"reflect"
)

func Dump(a ...interface{}) (s string) {
	s = ""
	for _, v := range a {
		s += "" + dumpSingle(v)
	}
	return
}

func dumpSingle(a interface{}) (s string) {
	v := reflect.Indirect(reflect.ValueOf(a))
	if v.Kind() == reflect.Invalid {
		return " (--) "
	}

	//v := reflect.ValueOf(a)
	if !v.IsValid() {
		return ""
	}
	switch v.Kind() {
	case reflect.String, reflect.Bool, reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		{
			s = fmt.Sprint(a)
			return
		}
	case
		reflect.Uintptr, reflect.UnsafePointer, reflect.Ptr:
		{
			if v.CanAddr() {
				p := reflect.Indirect(v.Addr())
				return Dump(p)
			}
			return ""
		}
	case reflect.Array, reflect.Slice:
		{
			s = ""
			for i := 0; i < v.Len(); i++ {
				if len(s) == 0 {
					s = Dump(v.Index(i).Interface())
				} else {
					s += "," + Dump(v.Index(i).Interface())
				}
			}
			return "[" + s + "]"
		}
	case reflect.Chan:
		{
			return "chan type can not dump to string "
		}
	case reflect.Map:
		{
			if len(v.MapKeys()) == 0 {
				return ""
			}

			it := v.MapRange()
			for it.Next() {
				key := it.Key()
				value := it.Value().Interface()
				if len(s) == 0 {
					s = fmt.Sprint(key, ":", Dump(value))
				} else {
					s += "," + fmt.Sprint(key, Dump(value))
				}
			}
			return "[" + s + "]"
		}
	case reflect.Struct:
		{
			return dumpStruct(v)
		}
	default:
		s = fmt.Sprint(v.Interface())
	}
	return
}

func dumpStruct(v reflect.Value) (s string) {
	if v.Kind() != reflect.Struct {
		return ""
	}

	s = ""
	for i := 0; i < v.Type().NumField(); i++ {
		fd := v.Type().Field(i)

		fdValue := v.FieldByName(fd.Name)
		if !fdValue.IsValid() {
			continue
		}

		sub := Dump(fdValue.Interface())
		if len(s) == 0 {
			s = fmt.Sprint(fd.Name, ":", sub)
		} else {
			s += "," + fmt.Sprint(fd.Name, ":", sub)
		}
	}

	return "{" + s + "}"
}
