package yred

import "github.com/vhaoran/vchat/common/typeconv"

func DoGet(key string) *typeconv.StrResult {
	s, err := X.Get(key).Result()
	if err != nil {
		return typeconv.NewStrResult("")
	}
	//----------------------------------------------
	return typeconv.NewStrResult(s)
}

func DoHGet(key, field string) *typeconv.StrResult {
	s, err := X.HGet(key, field).Result()
	if err != nil {
		return typeconv.NewStrResult("")
	}
	//----------------------------------------------
	return typeconv.NewStrResult(s)
}
