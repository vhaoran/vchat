package yqiniu

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type Abc struct {
	ID     int
	Name   string
	Url    string
	Remark string
}

type AbcD struct {
	Abc
	X string
}

func Test_struct_value(t *testing.T) {
	bean := &AbcD{
		Abc: Abc{
			ID:     1,
			Name:   "name_1",
			Url:    "this is url",
			Remark: "this is remark",
		},
	}
	err := Key2UrlOfQiNiu(bean, "Name", "Url", "Remark")
	fmt.Println(" ****  ", err)
	spew.Dump(bean)
}
