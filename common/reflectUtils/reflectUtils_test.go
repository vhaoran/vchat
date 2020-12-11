package reflectUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type Good struct {
	ID   int
	Name string
}

func Test_isPointer(t *testing.T) {
	a := new(Good)
	//
	log.Println(IsPointer(a))

	i := 5
	pi := &i
	fmt.Println("------pi" + "-----------")
	log.Println(IsPointer(pi))
	//
	m := map[int]int{1: 2, 2: 3}
	fmt.Println("------m" + "-----------")
	log.Println(IsPointer(m))

}

func Test_isSlice(t *testing.T) {
	l := make([]*Good, 0)
	fmt.Println("------", IsSlice(l), "-----------")
	fmt.Println("------", IsSlice(&l), "-----------")
}

func Test_make_alice_element_ptr(t *testing.T) {
	l := make([]*Good, 0)
	a, _ := MakeSliceElemPtr(l)
	spew.Dump(a)
}

func Test_MakeStructObj(t *testing.T) {
	bean := &Good{
		ID:   1,
		Name: "aaaaaaa",
	}
	buffer, err := json.Marshal(bean)
	if err != nil {
		log.Println(err)
		return
	}

	//
	obj, er1 := MakeStructObj(new(Good))
	if er1 != nil {
		log.Println(err)
		return
	}

	er2 := json.Unmarshal(buffer, obj)
	if er2 != nil {
		log.Println(er2)
		return
	}
	spew.Dump(obj)
}
