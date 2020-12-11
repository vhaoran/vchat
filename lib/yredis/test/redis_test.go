package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/lib/yredis"
)

type Good struct {
	ID   int64
	Name string
}

func (Good) TableName() string {
	return "good"
}

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		LoadRedis:        true,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_call_back_set(t *testing.T) {
	for i := 0; i < 100; i++ {
		t0 := time.Now()
		k := i % 10
		v, err := yredis.CacheAutoGetH(new(Good), int64(k),
			func() (interface{}, error) {
				time.Sleep(50 * time.Millisecond)

				return &Good{
					ID:   int64(k),
					Name: "name-" + fmt.Sprint(int64(k)),
				}, nil

			}, time.Second*100)
		fmt.Println("------", err, "-----------")
		spew.Dump(v)
		fmt.Println("------", time.Since(t0), "-----------")
	}
}

func Test_CacheClearH(t *testing.T) {
	yredis.CacheClearH(new(Good), 2, 4, 8)
}

func Test_debug(t *testing.T) {
	ylog.Debug("hello")
}

func Test_zz(t *testing.T) {
	yredis.CacheClearH(new(Good))
}

func Test_CacheGet_get(t *testing.T) {
	for i := 0; i < 5; i++ {
		l, err := yredis.CacheGet("aaa", i,
			time.Second*1000,
			func() (interface{}, error) {
				//如果 Key值不在缓存中时，如何得到成它
				return []string{fmt.Sprint(i),
					fmt.Sprint(i),
				}, nil
			},
			func(s string) (interface{}, error) {
				l := make([]string, 0)
				if err := json.Unmarshal([]byte(s), &l); err != nil {
					return nil, err
				}
				return l, nil
			},
		)
		ylog.Debug("--------redis_test.go------", l)
		ylog.Debug("--------redis_test.go------", err)
	}
}
