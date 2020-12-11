package test

import (
	"testing"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/yqiniu"
)

func init() {
	//------------ prepare modules----------
	//本步骤主要是装入系统必备的模块
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: true, //这不同必需要的
		LoadEtcd:         true, //etcd必須開啟，否則無法自動發現服務
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadJwt:          false,
		LoadQiniu:        true,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_get_token(t *testing.T) {
	z, err := yqiniu.GetToken(30)
	ylog.Debug("--------q_test.go------", z)
	ylog.Debug("--------q_test.go------", err)
}
