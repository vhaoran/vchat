package unit

import (
	"testing"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yetcd"
	"github.com/vhaoran/vchat/lib/ylog"
)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         true,
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_watch_etcd(t *testing.T) {
	_ = yetcd.RegisterService("/test", "a", "8888")

	ylog.Debug("-----------------------------")
	yetcd.InspectLoop("/test", "a", "8888")
}
