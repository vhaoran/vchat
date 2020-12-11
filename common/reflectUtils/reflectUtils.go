package reflectUtils

import (
	"errors"
	"log"
	"reflect"
)

//make a Point of slice element
//
func MakeSliceElemPtr(a interface{}) (interface{}, error) {
	if !IsSlice(a) {
		return nil, errors.New("input muse a slice")
	}
	//
	v := reflect.Indirect(reflect.ValueOf(a))
	//
	tp := v.Type().Elem()
	//
	if tp.Kind() == reflect.Ptr {
		log.Println(tp.Elem())
		return reflect.New(tp.Elem()).Interface(), nil
	}

	return reflect.New(tp).Interface(), nil
}

//传入一个interface型的结构或结构指针，返回新的结果对象
func MakeStructObj(ptr interface{}) (obj interface{}, err error) {
	bean := reflect.Indirect(reflect.ValueOf(ptr))
	//
	if bean.Type().Kind() != reflect.Struct {
		return nil, errors.New("no a struct")
	}

	//
	return reflect.New(bean.Type()).Interface(), nil
}
