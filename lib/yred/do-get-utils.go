package yred

import "github.com/vhaoran/vchat/common/typeconv"

func DoGet(key string) *typeconv.Str {
	s, err := X.Get(key).Result()
	if err != nil {
		return typeconv.NewStrData("")
	}
	//----------------------------------------------
	return typeconv.NewStrData(s)
}

func DoHGet(key, field string) *typeconv.Str {
	s, err := X.HGet(key, field).Result()
	if err != nil {
		return typeconv.NewStrData("")
	}
	//----------------------------------------------
	return typeconv.NewStrData(s)
}
