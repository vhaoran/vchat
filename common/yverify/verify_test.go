package yverify

import (
	"errors"
	"fmt"
	"testing"
)

type (
	Abc struct {
		A string
	}
)

func Test_aaa(t *testing.T) {
	var obj *Abc = &Abc{}

	if err := NewObj(false).
		Gt("fd1", 3, 20).
		Lt("fd2", 40, 5).
		Fn(errors.New("err2")).
		FnBool("fnBool", false).
		NotEmpty("my name", "").
		NotEmpty("my name", make(map[string]string)).
		NotEmpty("my name", make([]*struct{}, 0)).
		NotEmpty("my name", [1]string{""}).
		NotEmpty("obj", obj.A).
		//NotEmpty("obj", obj.A).
		Err(); err != nil {
		fmt.Println("ret: ", err.Error())
	}
}
