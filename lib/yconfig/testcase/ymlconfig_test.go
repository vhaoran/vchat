package testcase

import (
	"fmt"
	"github.com/vhaoran/vchat/lib/yconfig"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type testYml struct {
	A    int      `json:"a"`
	B    string   `json:"b"`
	List []string `json:"list"`
}

func Test_get_yml_config(t *testing.T) {
	bean, err := yconfig.GetYmlConfig()
	fmt.Println("-----------------", bean, err)
	fmt.Println("-----------------")
	spew.Dump(bean)
	fmt.Println("-----------------")
	spew.Dump(bean.ES)
	fmt.Println("-----------------")
	spew.Dump(bean.Qiniu)
	fmt.Println("-----------------")
	spew.Dump(bean.Bot)
	fmt.Println("-----------------")
}

func Test_cfg_2_get(t *testing.T) {
	//
	bean := new(testYml)
	err := yconfig.GetYmlConfigOfType(bean, "a_config")
	spew.Dump(err)
	spew.Dump(bean)
}
