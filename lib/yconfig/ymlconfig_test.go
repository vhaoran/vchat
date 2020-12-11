package yconfig

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_get_yml_config(t *testing.T) {
	bean, err := GetYmlConfig()
	fmt.Println("-----------------", bean, err)
	fmt.Println("-----------------")
	spew.Dump(bean)
	fmt.Println("-----------------")
	spew.Dump(bean.ES)
	fmt.Println("-----------------")
	spew.Dump(bean.Qiniu)
	fmt.Println("-----------------")

}
