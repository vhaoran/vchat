package g

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

type (
	Good struct {
		ID   int64
		Name string
	}
)

func Test_get_buffer(t *testing.T) {
	s, _ := GetBufferForMq("abc")
	log.Println(string(s))

	s, _ = GetBufferForMq([]byte("abc"))
	log.Println(string(s))

	s, _ = GetBufferForMq([]int{1, 2, 3})
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq(3)
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq(Good{ID: 1, Name: "33"})
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq([]Good{Good{ID: 1, Name: "33"}, Good{ID: 2, Name: "3344"}})
	fmt.Println("------", string(s), "-----------")
}

func Test_exec_path(t *testing.T) {
	p, err := GetExecPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("----path:------", p, "------------")
}

func Test_path_exists(t *testing.T) {
	l := []string{
		"~/a",
		"~/go/src",
		"./",
		"/usr",
		"/etc/",
	}

	for _, p := range l {
		ok, err := PathExists(p)
		fmt.Println("---", p)
		fmt.Println("-----------", ok)
		fmt.Println("-----------", err)
	}

}

func Test_cn2tra(t *testing.T) {
	s := "迪拜（阿拉伯语：دبي，英语：Dubai），是阿拉伯联合酋长国人口最多的城市，位于波斯湾东南海岸，迪拜也是组成阿联酋七个酋长国之一——迪拜酋长国的首都。"
	str, err := cn2tradiation(s)
	spew.Dump(err)
	spew.Dump(str)
}
