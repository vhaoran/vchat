package ymongo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/lib/ylog"
)

func GetMapOfBsonD(i interface{}) (ofMap interface{}) {
	if reflectUtils.IsNil(i) {
		return nil
	}

	//--------bsonD -----------------------------
	if isMongoBsonD(i) {
		return fromBsonD(i)
	}

	//--------map -----------------------------
	if reflectUtils.IsMap(i) {
		m := make(map[string]interface{})

		v := reflect.Indirect(reflect.ValueOf(i))
		//
		l := v.MapKeys()
		for _, key := range l {
			value := v.MapIndex(key)
			//
			bsonV := GetMapOfBsonD(value.Interface())
			//v.SetMapIndex(key, reflect.ValueOf(bsonV))
			m[key.String()] = bsonV
		}
		ofMap = m
		return
	}

	//------array ------------------------------
	if isMongoBsonDSlice(i) {
		l := make([]interface{}, 0)
		//
		v := reflect.Indirect(reflect.ValueOf(i))
		for ii := 0; ii < v.Len(); ii++ {
			src := v.Index(ii)
			bean := GetMapOfBsonD(src.Interface())
			l = append(l, bean)
		}
		return l
	}

	//---others-------
	return i
}

func fromBsonD(i interface{}) (ofMap interface{}) {
	ofMap = i

	//if reflectUtils.IsSlice(i) {
	//	v := reflect.ValueOf(i)
	//	for i := 0; i < v.Len(); i++ {
	//		vv := v.Index(i)
	//		vv.Set(reflect.ValueOf(GetMapOfBsonD(vv.Interface())))
	//	}
	//	return v.Interface()
	//}

	if !isMongoBsonD(i) {
		return
	}

	//error handle
	defer func() {
		if err := recover(); err != nil {
			ylog.Error("tran-d-map.go->", err)
			ofMap = i
		}
	}()

	v := reflect.Indirect(reflect.ValueOf(i))
	b := v.Interface().(bson.D)

	m := make(map[string]interface{})
	for _, each := range b {
		m[each.Key] = each.Value
	}

	//------------递归搜索设置------------
	for k, vv := range m {
		m[k] = GetMapOfBsonD(vv)
	}

	ofMap = m
	return
}

func isMongoBsonDSlice(i interface{}) bool {
	if !reflectUtils.IsSlice(i) {
		return false
	}
	v := reflect.Indirect(reflect.ValueOf(i))
	for ii := 0; ii < v.Len(); ii++ {
		each := v.Index(ii)
		if isMongoBsonD(each.Interface()) {
			return true
		}
	}
	return false
}

func isMongoBsonD(i interface{}) bool {
	switch i.(type) {
	case bson.D:
		return true
	case *bson.D:
		return true
	default:
		return false
	}
}
