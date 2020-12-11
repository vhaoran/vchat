package main

import (
	"context"
	golog "log"
	"net/http"

	"github.com/vhaoran/vchat/lib/ykit"
	//单独运行时导入改为这个
	// or import "github.com/vhaoran/vchat"
	"github.com/vhaoran/vchat/demo/unit/intf"

	"github.com/vhaoran/vchat/lib"
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

//gateway功能不需要每一个模块来实现，但用这个模块可以测试微服务是否能补成功调用
func main() {

	addr := "localhost:9999"
	//ctx := context.Background()
	mux := http.NewServeMux()

	mux.Handle("/api/CtxTest", new(intf.CtxTestHandler).HandlerSD(nil))
	mux.Handle("/api/MapTest", new(intf.MapTestH).HandlerSD(nil))

	mux.Handle("/api/MapTestCommon", new(ykit.RootTran).HandlerSDCommon(
		context.Background(),
		"api",
		"POST",
		"/MapTest",
		nil))

	mux.Handle("/api/pb",
		new(ykit.RootTran).HandlerSDCommon(
			context.Background(),
			"api",
			"POST",
			"/pb",
			nil))

	golog.Println(
		`start at :9999,url is curl:localhost/hello`,
		`test command:`,
		`curl -X POST -H 'Authorization:whr_token' -H 'Content-Type:application/json'  -d '{"S":"hello,world pass in data"}' localhost:9999/api/CtxTest -v`)

	golog.Fatal(http.ListenAndServe(addr, mux))
}
