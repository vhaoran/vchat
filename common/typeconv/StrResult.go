package typeconv

import (
	"encoding/json"
	"errors"
	"github.com/vhaoran/vchat/common/reflectUtils"
	"strconv"
	"strings"
)

type StrData struct {
	Text string
}

func NewStrData(s string) *StrData {
	return &StrData{
		Text: s,
	}
}

//只需要传入一个，没有时默认为0
func (r *StrData) AsInt64(defValues ...int64) int64 {
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

func (r *StrData) AsInt(defValues ...int) int {
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

func (r *StrData) AsStr() string {
	return r.Text
}
func (r *StrData) Str() string {
	return r.Text
}

func (r *StrData) AsBool(defValues ...bool) bool {
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
func (r *StrData) Unmarshal(ptr interface{}) error {
	//不是指针是原样返回
	if !reflectUtils.IsPointer(ptr) {
		return errors.New("passed param must be a pointer")
	}
	//----------------------------------------------
	return json.Unmarshal([]byte(r.Text), ptr)
}
