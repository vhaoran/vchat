package yes

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/lib/yconfig"
)

func Test_es_cnt(t *testing.T) {
	cfg := yconfig.ESConfig{
		Url:   []string{"http://192.168.0.99:9200"},
		Sniff: false,
	}

	c, err := NewESCnt(cfg)
	fmt.Println("------ok----------")
	spew.Dump(err)
	spew.Dump(c)
}
