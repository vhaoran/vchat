package ymodel

import (
	"reflect"
)

func TableName(ptr interface{}) string {
	//
	if ptr == nil {
		return ""
	}

	bean := reflect.Indirect(reflect.ValueOf(ptr))
	if bean.Type().Kind() != reflect.Struct {
		return ""
	}

	//
	fun := bean.MethodByName("TableName")
	if !fun.IsValid() {
		return ""
	}
	l := fun.Call(nil)
	if len(l) > 0 {
		return l[0].String()
	}
	return ""
}
