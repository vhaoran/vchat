package ymodel

import (
	"fmt"
	"testing"
)

type Abc struct {
	ID   int
	Name string
}

func (Abc) TableName() string {
	return "abc"
}

func Test_get_tb_name(t *testing.T) {
	s := TableName(new(Abc))
	fmt.Println("------", s, "-----------")
}

