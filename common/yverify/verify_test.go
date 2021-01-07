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

func Test_chain_call(t *testing.T) {
	a := func() error {
		fmt.Println(" a called ")
		return nil
	}
	b := func() error {
		fmt.Println(" b called ")
		return nil
	}
	c := func() error {
		fmt.Println(" c called ")
		return nil
	}

	if err := NewObj().
		Fn(a()).
		Fn(b()).
		Fn(c()).
		Err(); err != nil {
		fmt.Println("########## error: ", err)
	}
	fmt.Println("------ok----")
}

func Test_fn_call_a(t *testing.T) {
	a := func() (string, error) {
		fmt.Println(" a called ")
		return "", nil
	}
	b := func() error {
		fmt.Println(" b called ")
		return nil
	}
	c := func() (int, error) {
		fmt.Println(" c called ")
		return 1, nil
	}

	if err := NewObj().
		Fn(func() error {
			_, err := a()
			return err
		}()).
		Fn(func() error {
			return b()
		}()).
		Fn(func() error {
			_, err := c()
			return err
		}()).
		Err(); err != nil {
		fmt.Println("########## error: ", err)
	}

	fmt.Println("------ok----")
}
