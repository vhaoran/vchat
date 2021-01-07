package ypage

import (
	"fmt"
	"github.com/vhaoran/vchat/common/reflectUtils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"reflect"
	"strings"
)

type SqlWhere struct {
}

func (r *SqlWhere) GetWhere(a bson.D) (exp string, params []interface{}) {
	if a == nil {
		return "1 = 1", nil
	}
	//
	return r.scan("", a)
}

func (r *SqlWhere) isOr(sign string) bool {
	return sign == "$or"
}

func (r *SqlWhere) isSignConn(s string) bool {
	return strings.Contains(s, "$or") ||
		strings.Contains(s, "$and")
}

func (r *SqlWhere) isSignLogic(s string) bool {
	return len(r.getSign(s)) > 0
}

func (r *SqlWhere) scan(sign string, m bson.D) (exp string, params []interface{}) {
	exp = ""
	params = make([]interface{}, 0)

	signSQL := " and "
	if r.isOr(sign) {
		signSQL = " or "
	}

	for _, v := range m {
		str, p := r.getE(v)
		if len(exp) == 0 {
			exp = str
		} else {
			exp += signSQL + str
		}

		params = append(params, p...)
	}

	return
}

func (r *SqlWhere) getSign(key string) string {
	signMap := map[string]string{
		"$gt":  " > ",
		"$gte": " >= ",
		"$lt":  " < ",
		"$lte": " <= ",
		"$in":  " in ",

		"$ne":  " != ",
		"$nin": " not in ",
	}

	s, ok := signMap[key]
	if ok {
		return s
	}
	return ""
}

func (r *SqlWhere) getE(m primitive.E) (exp string, p []interface{}) {
	p = make([]interface{}, 0)

	k := m.Key
	v := m.Value

	//--------key------------------
	// is and / or
	if r.isSignConn(k) {
		//must be bson.D
		d := r.Slice2D(v)
		//
		return r.scan(k, d)
	}

	// is > >= < <= in
	if r.isSignLogic(k) {
		//must be bson.D
		exp = r.getSign(k) + " (?) "
		p = append(p, v)
		return
	}

	//--------------value-----------------------------------
	if r.IsE(v) {
		exp1, p1 := r.getE(v.(primitive.E))
		exp = k + exp1
		p = append(p, p1...)
		return
	}

	//
	if r.IsMap(v) {
		//--------is map -----------------------------
		es := r.Map2E(v)
		if len(es) > 0 {
			exp1, p1 := r.getE(es[0])
			exp = fmt.Sprintf(" %s %s ", k, exp1)
			p = append(p, p1...)
			return
		}
		return
	}

	exp = fmt.Sprintf(" %s = ?", k)
	p = append(p, v)
	return
}

func (r *SqlWhere) Map2E(i interface{}) []bson.E {
	if i == nil {
		return nil
	}

	l := make([]bson.E, 0)

	//-------- -----------------------------
	v := reflect.Indirect(reflect.ValueOf(i))
	for _, fd := range v.MapKeys() {
		value := v.MapIndex(fd)
		fmt.Println(fd, value)

		bean := bson.E{
		}
		bean.Key = fd.String()
		bean.Value = value.Interface()

		l = append(l, bean)
	}

	return l
}

func (r *SqlWhere) IsE(i interface{}) bool {
	if i == nil {
		return false
	}

	switch i.(type) {
	case primitive.E:
		return true
	}

	return false
}

func (r *SqlWhere) IsMap(i interface{}) bool {
	if i == nil {
		return false
	}

	//-------- -----------------------------
	tp := reflect.TypeOf(i)
	switch tp.Kind() {
	case reflect.Map:
		return true
	default:
		log.Println("tp.kind:", tp.Kind())
	}

	v := reflect.Indirect(reflect.ValueOf(i))
	switch v.Kind() {
	case reflect.Map:
		return true
	case reflect.Struct:
		//
		fmt.Println("vTYpe():", v.Type())
		fmt.Println("vTYpe().String():", v.Type().String())

	default:
		log.Println("v.kind:", v.Kind())
	}

	return false
}

func (r *SqlWhere) Slice2D(i interface{}) bson.D {
	if i == nil {
		return nil
	}

	if !reflectUtils.IsSlice(i) {
		return nil
	}
	v := reflect.Indirect(reflect.ValueOf(i))

	l := bson.D{}
	//-------- -----------------------------
	//l = append(l, nil)
	for i := 0; i < v.Len(); i++ {
		m := v.Index(i).Interface()
		if r.IsE(m) {
			l = append(l, m.(primitive.E))
			continue
		}

		if r.IsMap(m) {
			//
			e := r.Map2E(m)
			//
			l = append(l, e...)
		}
	}

	return l
}
