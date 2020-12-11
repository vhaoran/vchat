package yqiniu

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/lib/ylog"
)

func Key2UrlOfQiNiu(obj interface{}, fields ...string) error {
	defer func() {
		if err := recover(); err != nil {
			ylog.Debug("--------tran-image-url.go------", err)
		}
	}()

	if !reflectUtils.IsStruct(obj) {
		return errors.New("不是结构")
	}
	if !reflectUtils.IsPointer(obj) {
		return errors.New("不是指针，无法继续")
	}

	if reflect.ValueOf(obj).IsNil() {
		return errors.New("必须是指针，才能操作")
	}

	v := reflect.Indirect(reflect.ValueOf(obj))

	//
	for _, fdName := range fields {
		if _, ok := v.Type().FieldByName(fdName); ok {
			//i := fd.Index
			key := v.FieldByName(fdName).String()
			fmt.Println(" --- fdValue:", key)
			url := GetVisitURL(key)
			//
			fmt.Println(" --- url:", url)
			//
			v.FieldByName(fdName).SetString(url)
			continue
		}
	}
	return nil
}
