package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/ykit"
	"github.com/vhaoran/vchat/lib/ylog"
)

//elasticSearch

func Test_add_single(t *testing.T) {
	bean := Product{
		ID:       "1",
		Name:     fmt.Sprint("name_3"),
		CateID:   0,
		CateName: fmt.Sprint("cate_", 1),
		Tag:      "汽车 飞机 大炮",
		Price:    float32(10.0),
		Remark:   "this is a good test",
	}


	r, err := yes.X.Index().
		Index("t").
		Id("111a").
		BodyJson(bean).Do(context.Background())
	ylog.Debug("--------yes_search_test.go------", err)
	ylog.DebugDump("--------yes_search_test.go------", r)
	//

}

func Test_update_rec(t *testing.T) {
	r, err := yes.X.Update().Index("t").
		Id("111").
		Doc(ykit.M{
			"tag": "test value tag1 tag2",
		}).
		Do(context.Background())

	ylog.DebugDump("--------r------", r)
	ylog.Debug("-------err------", err)

	ret, err1 := yes.X.Get().Index("t").
		Id("111").
		Do(context.Background())

	ylog.DebugDump("--------------", ret)
	ylog.Debug("--------r------", err1)
	s, _ := ret.Source.MarshalJSON()
	ylog.Debug("--------yes_crud_test.go------", string(s))
}


