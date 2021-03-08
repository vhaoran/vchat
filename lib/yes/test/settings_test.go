package test

import (
	"context"
	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/ylog"
	"testing"
)

func Test_settings(t *testing.T) {
	body := `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
                "retweets":{
                    "type":"long"
                },
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
	}
}
`
	// Put settings
	r, err := yes.X.CreateIndex("abc").Body(body).Do(context.Background())
	ylog.Debug("--------yes_search_test.go--- ----", err)
	ylog.DebugDump("--------yes_search_test.go--- ----", r)
}
