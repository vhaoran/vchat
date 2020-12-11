package yverify

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/common/reflectUtils"
)

type (
	VerifyOBJ struct {
		onErrStop bool
		Errs      []string
	}
)

//输入为整形值 或浮点娄，必须大于0
func NewObj(onErrStop ...bool) *VerifyOBJ {
	b := true
	if len(onErrStop) > 0 {
		b = onErrStop[0]
	}

	return &VerifyOBJ{
		onErrStop: b,
		Errs:      make([]string, 0),
	}
}

func (r *VerifyOBJ) hasErr() bool {
	return r.Errs != nil && len(r.Errs) > 0
}

func (r *VerifyOBJ) needContinue() bool {
	return !r.hasErr() || (r.hasErr() && !r.onErrStop)
}

func (r *VerifyOBJ) NotZero(name string, src interface{}) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsInt(src) {
		return r
	}
	//
	if ii, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		if ii == 0 {
			msg := fmt.Sprintf("[%s]值[%d]必须不为0", name, ii)

			r.push(msg)
			return r
		}
	}
	return r
}

func (r *VerifyOBJ) GtF(name string, src interface{}, l ...float64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := float64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		f2 := dst
		if !(f1 > f2) {
			msg := fmt.Sprintf("[%s]值[%f]必须大于[%f]", name, f1, f2)

			r.push(msg)
			return r
		}
	}
	return r
}

func (r *VerifyOBJ) Gt(name string, src interface{}, l ...int64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := int64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if i1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		i2 := dst
		if !(i1 > i2) {
			msg := fmt.Sprintf("[%s]值[%d]必须大于[%d]", name, i1, i2)
			r.push(msg)
			return r
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) GteF(name string, src interface{}, l ...float64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := float64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		f2 := dst
		if !(f1 >= f2) {
			msg := fmt.Sprintf("%s 值[%f]必须大于或等于[%f]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) Gte(name string, src interface{}, l ...int64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := int64(0)
	if len(l) > 0 {
		dst = l[0]
	}
	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		f2 := dst
		if !(f1 >= f2) {
			msg := fmt.Sprintf("[%s]值[%d]必须大于或等于[%d]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

//输入为整形值 或浮点，src<dst
func (r *VerifyOBJ) LtF(name string, src interface{}, l ...float64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := float64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		f2 := dst
		if !(f1 < f2) {
			msg := fmt.Sprintf("[%s]值[%f]必须小于[%f]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

//输入为整形值 或浮点，src<dst
func (r *VerifyOBJ) Lt(name string, src interface{}, l ...int64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := int64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		f2 := dst
		if !(f1 < f2) {
			msg := fmt.Sprintf("[%s]值[%d]必须小于[%d]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

//>= l && <= h
func (r *VerifyOBJ) Between(name string, src, low, high interface{}) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(low) || !reflectUtils.IsNumber(high) {
		return r
	}

	i, _ := strconv.ParseInt(fmt.Sprint(src), 10, 64)
	l, _ := strconv.ParseInt(fmt.Sprint(low), 10, 64)
	h, _ := strconv.ParseInt(fmt.Sprint(high), 10, 64)
	if src == nil {
		return r
	}

	//
	if !(i >= l && i <= h) {
		msg := fmt.Sprintf("[%s]值[%d]必须介于[%d]和[%d]之间", name, i, l, h)
		r.push(msg)
	}

	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) LteF(name string, src interface{}, l ...float64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := float64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseFloat(fmt.Sprint(src), 64); err == nil {
		f2 := dst
		if !(f1 <= f2) {
			msg := fmt.Sprintf("[%s]值[%f]必须小于或等于[%f]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

//输入为整形值 或浮点娄，必须大于0
func (r *VerifyOBJ) Lte(name string, src interface{}, l ...int64) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	dst := int64(0)
	if len(l) > 0 {
		dst = l[0]
	}

	if r.hasErr() && r.onErrStop {
		return r
	}

	if src == nil {
		return r
	}
	//
	if !reflectUtils.IsNumber(src) || !reflectUtils.IsNumber(dst) {
		return r
	}
	//
	if f1, err := strconv.ParseInt(fmt.Sprint(src), 10, 64); err == nil {
		f2 := dst
		if !(f1 <= f2) {
			msg := fmt.Sprintf("[%s]值[%d]必须小于或等于[%d]", name, f1, f2)
			r.push(msg)
			return r
		}
	}
	return r
}

func (r *VerifyOBJ) push(s string) {
	if r.Errs == nil {
		r.Errs = make([]string, 0)
	}
	r.Errs = append(r.Errs, s)
}

//判断鞭值是否为空指针
func (r *VerifyOBJ) NotNilPtr(name string, l ...interface{}) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if l == nil {
		msg := fmt.Sprintf("[%s]不能为空", name)
		r.push(msg)
		return r
	}

	for _, ptr := range l {
		if !reflectUtils.IsPointer(ptr) {
			continue
		}
		//
		v := reflect.ValueOf(ptr)
		if v.IsNil() {
			msg := fmt.Sprintf("[%s]不能为空", name)
			r.push(msg)
			return r
		}
	}
	return r
}

// l is string,array,slice,amp
func (r *VerifyOBJ) NotEmpty(name string, lst ...interface{}) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if lst == nil {
		msg := fmt.Sprintf("[%s]不能为空", name)
		r.push(msg)
		return r
	}

	for _, l := range lst {
		if l == nil {
			msg := fmt.Sprintf("[%s]不能为空", name)
			r.push(msg)
			return r
		}
		if reflectUtils.IsPointer(l) {
			if reflectUtils.IsNil(l) {
				msg := fmt.Sprintf("[%s]不能为空,长度必须大于0", name)
				r.push(msg)
				return r
			}
		}

		v := reflect.Indirect(reflect.ValueOf(l))
		i := 1
		switch v.Kind() {
		case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
			i = v.Len()
		default:
			i = 1
		}

		if i <= 0 {
			msg := fmt.Sprintf("[%s]不能为空,长度必须大于0", name)
			r.push(msg)
			return r
		}
	}
	return r
}

func (r *VerifyOBJ) InSlice(name string, src, l interface{}) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if g.InSlice(src, l) {
		return r
	}
	msg := fmt.Sprintf("[%s]元素不在列表中", name)
	r.push(msg)
	return r
}

//这是自定义的fn，用户可以传入自定义验证结果
func (r *VerifyOBJ) Fn(l ...error) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	for _, err := range l {
		if err != nil {
			r.push(err.Error())
		}
	}
	return r
}

func (r *VerifyOBJ) FnTailErr(_ interface{}, err error) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if err != nil {
		r.push(err.Error())
	}
	return r
}

func (r *VerifyOBJ) FnTrue(name string, b bool) *VerifyOBJ {
	return r.FnBool(name, b)
}

func (r *VerifyOBJ) FnBool(name string, b bool) *VerifyOBJ {
	if !r.needContinue() {
		return r
	}

	if !b {
		r.push(fmt.Sprintf("验证失败，[%s] ", name))
	}
	return r
}

//这是链式语法验证的结果
func (r *VerifyOBJ) Err() error {
	if len(r.Errs) > 0 {
		s := strings.Join(r.Errs, ";  ")
		return errors.New(s)
	}
	return nil
}
