package typeconv

import (
	"encoding/json"
	"errors"
	"github.com/vhaoran/vchat/common/reflectUtils"
	"strconv"
	"strings"
)

type StrResult struct {
	Text string
}

func NewStrResult(s string) *StrResult {
	return &StrResult{
		Text: s,
	}
}

//只需要传入一个，没有时默认为0
func (r *StrResult) AsInt64(defValues ...int64) int64 {
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

func (r *StrResult) AsInt(defValues ...int) int {
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

func (r *StrResult) AsStr() string {
	return r.Text
}
func (r *StrResult) Str() string {
	return r.Text
}

func (r *StrResult) AsBool(defValues ...bool) bool {
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
func (r *StrResult) Unmarshal(ptr interface{}) error {
	//不是指针是原样返回
	if !reflectUtils.IsPointer(ptr) {
		return errors.New("passed param must be a pointer")
	}
	//----------------------------------------------
	return json.Unmarshal([]byte(r.Text), ptr)
}
