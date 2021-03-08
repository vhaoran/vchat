package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	es "github.com/olivere/elastic/v7"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/ylog"
)

type Product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	CateID   int64  `json:"cate_id,omitempty"`
	CateName string `json:"cate_name,omitempty"`
	Tag      string `json:"tag,omitempty"`

	Price  float32 `json:"price,omitempty"`
	Remark string  `json:"remark,omitempty"`
}

func init() {
	_, err := lib.InitModulesOfOptions(
		&lib.LoadOption{
			LoadMicroService: false,
			LoadEtcd:         false,
			LoadPg:           false,
			LoadRedis:        false,
			LoadMongo:        false,
			LoadMq:           false,
			LoadRabbitMq:     false,
			LoadJwt:          false,
			LoadES:           true,
			LoadQiniu:        false,
		})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(" ---------- init ok---------")
}

func Test_add_batch(t *testing.T) {
	for i := int64(0); i < 100; i++ {
		bean := Product{
			//ID:       i,
			Name:     fmt.Sprint("name_", i),
			CateID:   0,
			CateName: fmt.Sprint("cate_", 1),
			Tag:      "汽车 飞机 大炮",
			Price:    float32(i * 10.0),
			Remark:   "this is a good test",
		}
		r, err := yes.X.Index().Index("index").BodyJson(bean).Do(context.Background())
		ylog.Debug("--------yes_search_test.go------", err)
		ylog.DebugDump("--------yes_search_test.go------", r)
	}
}

func Test_add_n(t *testing.T) {
	bean := Product{
		//ID:       i,
		Name:     fmt.Sprint("name_", time.Now().UnixNano()),
		CateID:   0,
		CateName: fmt.Sprint("cate_", 1),
		Tag:      "汽车 蓝天 白云",
		Price:    10.0,
		Remark:   "中国 魏浩然  李明 王伟",
	}
	r, err := yes.X.Index().Index("index").BodyJson(bean).Do(context.Background())
	ylog.Debug("--------yes_search_test.go------", err)
	ylog.DebugDump("--------yes_search_test.go------", r)
}

func Test_term(t *testing.T) {
	q := es.NewTermQuery("tag", "汽车")

	r, err := yes.X.Search("index").
		Query(q).
		Do(context.Background())
	ylog.Debug("--------yes_search_test.go------", err)
	ylog.DebugDump("--------yes_search_test.go------", r)
	output(r)
}

func Test_match_all(t *testing.T) {
	//
	//"query":{
	//	"match_all":{}
	//}
	q := es.NewRawStringQuery(`
    {
		"match_all":{}
    }
    `)

	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}
func Test_match_field_single(t *testing.T) {
	/*
		"match" : {
		         "city":"pune"
		      }
	*/
	q := es.NewRawStringQuery(`
    {
	"match" : {
	         "tag":"汽车"
	      }	
    }
    `)

	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

//
func Test_match_raw_str_query(t *testing.T) {
	q := es.NewRawStringQuery(`
      {
              "match": {
                 "tag": "飞机"
              }
      }
    `)
	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

func Test_match_multi_valueOfOneField_or(t *testing.T) {
	q := es.NewRawStringQuery(`
      {
		"bool": {
		  "should": [
			{ "match": { "tag": "飞机" }},
			{ "match": { "tag": "蓝天" }},
			{ "match": { "tag": "白云" }}
		  ]
		}                 
					 
      }
    `)
	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

func Test_match_multi_valueOfOneField_all(t *testing.T) {
	q := es.NewRawStringQuery(`
      {
        "query_string" : {
            "default_field" : "tag",
            "query" : "(白云) and (飞机)",
            "minimum_should_match": 2
        }					 
      }
    `)
	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

func Test_match_multi_field_multi_valueOfOneField_all(t *testing.T) {
	q := es.NewRawStringQuery(`
      {
        "query_string" : {
           "fields": [
                "tag",
                "remark"
            ],
            "query" : "(白云) and (飞机) and (魏浩然)",
            "minimum_should_match": 3
        }					 
      }
    `)
	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

func Test_match_multi_field_multi_valueOfOneField_all_qStr(t *testing.T) {
	q := es.NewQueryStringQuery("(白云) and (飞机) and (魏浩然)").Field("tag").
		Field("remark").
		MinimumShouldMatch("3")

	r, err := yes.X.Search("index").Query(q).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	ylog.DebugDump("----result---", r)
	output(r)
}

func output(r *es.SearchResult) {
	fmt.Println("-------- -----------------------------")
	for _, v := range r.Hits.Hits {
		str, err := json.Marshal(v)
		if err == nil {
			fmt.Println("##### :  ", string(str))
		} else {
			fmt.Println("***err:", err.Error())
		}
	}
}

