package typeconv

import (
	"encoding/json"
	"errors"
	"github.com/vhaoran/vchat/common/reflectUtils"
	"strconv"
	"strings"
)

type Str struct {
	Text string
}

func NewStrData(s string) *Str {
	return &Str{
		Text: s,
	}
}

func NewStr(s string) *Str {
	return &Str{
		Text: s,
	}
}

//只需要传入一个，没有时默认为0
func (r *Str) AsInt64(defValues ...int64) int64 {
	i0 := int64(0)
	if len(defValues) > 0 {
		i0 = defValues[0]
	}
	//----------------------------------------------
	i, err := strconv.ParseInt(r.Text, 10, 64)
	if err != nil {
		return i0
	}
	return i
}

//只需要传入一个，没有时默认为0
func (r *Str) AsFloat64(defValues ...float64) float64 {
	i0 := float64(0)
	if len(defValues) > 0 {
		i0 = defValues[0]
	}
	//----------------------------------------------
	i, err := strconv.ParseFloat(r.Text, 64)
	if err != nil {
		return i0
	}
	return i
}

func (r *Str) AsInt(defValues ...int) int {
	i0 := int(0)
	if len(defValues) > 0 {
		i0 = defValues[0]
	}
	//----------------------------------------------
	i, err := strconv.Atoi(r.Text)
	if err != nil {
		return i0
	}
	return i
}

func (r *Str) AsStr() string {
	return r.Text
}
func (r *Str) Str() string {
	return r.Text
}

func (r *Str) AsBool(defValues ...bool) bool {
	b := false
	if len(defValues) > 0 {
		b = defValues[0]
	}

	s := strings.ToLower(strings.Trim(r.Text, " "))
	switch s {
	case "":
		return b
	case "1":
		return true
	case "true":
		return true
	case "ok":
		return true
	case "yes":
		return true
	case "t":
		return true
	default:
		return false
	}
}

//from json str to object
func (r *Str) Unmarshal(ptr interface{}) error {
	//不是指针是原样返回
	if !reflectUtils.IsPointer(ptr) {
		return errors.New("passed param must be a pointer")
	}
	//----------------------------------------------
	return json.Unmarshal([]byte(r.Text), ptr)
}

func (r *Str) Replace(src, dst string) *Str {
	str := r.Text
	str = strings.Replace(str, src, dst, -1)
	return NewStrData(str)
}

func (r *Str) ReplaceMult(m map[string]string) *Str {
	str := r.Text
	for k, v := range m {
		str = strings.Replace(str, k, v, -1)
	}
	return NewStrData(str)
}

func (r *Str) RmSign(signList ...string) *Str {
	str := r.Text
	for _, v := range signList {
		str = strings.Replace(str, v, "", -1)
	}
	return NewStrData(str)
}

func (r *Str) ContainsAny(multStr ...string) bool {
	for _, v := range multStr {
		if strings.Contains(r.Text, v) {
			return true
		}
	}
	return false
}

func (r *Str) ContainsAll(multStr ...string) bool {
	for _, v := range multStr {
		if !strings.Contains(r.Text, v) {
			return false
		}
	}
	return true
}

//中间用seg分割
func (r *Str) RepeatN(n int, segs ...string) *Str {
	s := r.Text
	str := ""
	if len(segs) > 0 {
		str = segs[0]
	}

	//----------------------------------------------
	for i := 0; i < n; i++ {
		s += str + r.Text
	}
	return NewStr(s)
}
