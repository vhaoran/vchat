package test

import (
	"context"
	es "github.com/olivere/elastic/v7"
	"github.com/vhaoran/vchat/common/typeconv"
	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/ylog"
	"testing"
)

type Item struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func Test_settings(t *testing.T) {
	// ik_smart
	// ik_max_word
	body := `
{
	"settings": {
		"index": {
		  "analysis": {
			"analyzer":"ik_max_word"
		  }
		},
		"similarity": {
		  "my_bm25": {
			"type": "BM25",
			"k1": 0.01,
			"b": 1,
			"discount_overlaps": true
		  }
		}
	  },
	"mappings":{
			"properties":{
				"title":{
					"type":"text",
					"similarity":"my_bm25",
                    "analyzer": "ik_max_word",
                    "search_analyzer": "ik_smart"
				},
                "id":{
                    "type":"long",
                    "index": false
                }
			}
	}
}
`
	// Put settings
	_, _ = yes.X.DeleteIndex("abc").Do(context.Background())
	//
	r, err := yes.X.CreateIndex("abc").Body(body).Do(context.Background())
	//_, _ := yes.X.Index("abc").Body(body).Do(context.Background())
	ylog.Debug("--------yes_search_test.go--- ----", err)

	for i := 0; i < 3; i++ {
		bean := &Item{
			ID:    int64(i),
			Title: typeconv.NewStrData("微信").RepeatN(i, "good man ").Str(),
		}
		ret, err := yes.X.Index().Index("abc").BodyJson(bean).Do(context.Background())
		ylog.Debug("--------settings_test.go--- ----", err)
		ylog.DebugDump("--------settings_test.go--- ----", ret)
	}

	ylog.DebugDump("--------yes_search_test.go--- ----", r)
}

func Test_find(t *testing.T) {
	q := es.NewRawStringQuery(`
    {
	"match" : {
	         "title":"微信"
	      }	
    }
    `)

	r, err := yes.X.Search("abc").Query(q).Explain(true).
		Do(context.Background())
	ylog.Debug("-----err------", err)
	//ylog.DebugDump("----result---", r)
	output(r)
}
