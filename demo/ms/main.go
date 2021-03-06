package main

import (
	"fmt"
	golog "log"
	"net/http"

	"github.com/vhaoran/vchat/lib/ycaptcha"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yetcd"
	"github.com/vhaoran/vchat/lib/ylog"

	"github.com/vhaoran/vchat/demo/unit/impl"
	"github.com/vhaoran/vchat/demo/unit/intf"
)

var (
	// 每个微服务都不同，这里需要更改
	// todo
	msTag   = "api"
	host    = "0.0.0.0"
	port    = 9000
	regHost = "127.0.0.1"
	regPort = 9000
)

func init() {
	//------------ prepare modules----------
	//本步骤主要是装入系统必备的模块
	cfg, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadEtcd:  true,
		LoadPg:    false,
		LoadRedis: true,
		LoadMongo: false,
		LoadMq:    false,
		LoadJwt:   true,
		LoadQiniu: true,
	})
	if err != nil {
		panic(err.Error())
	}

	//-------###############-----------------------------------
	//装入微服务配置,对于微服务开发人员，中需要写跌幅，写接口及实现
	ms := cfg.MicroService
	//assert yconfig.config
	if g.IsEmptyOr(ms.Host, ms.Tag, ms.RegHost) {
		panic("配置文件错误，必须有microService.host/regHost/tag")
	}

	if g.IsZeroOr(ms.Port, ms.RegPort) {
		panic("配置文件错误，必须有microService.port/regPort")
	}

	//微服務tag,用於註冊時標識/監聽端口/監聽主機
	msTag, port, regPort = ms.Tag, ms.Port, ms.RegPort
	//注冊用監聽/注册用主机
	host, regHost = "0.0.0.0", ms.RegHost

	ycaptcha.Init()
}

func main() {
	ylog.Info("微服务tag：", msTag)

	mux := http.NewServeMux()
	//--------handlers-----------------------------
	// 每一步：配置路由
	// HelloWorld handler

	// 每一个微服务都需要实现的方法，用于测试服务是否运行
	mux.Handle("/ping", http.HandlerFunc(new(Ping).handler))
	mux.Handle("/CtxTest", new(intf.CtxTestHandler).HandlerLocal(new(ctl.CtxTestImpl), nil))
	mux.Handle("/MapTest", new(intf.MapTestH).HandlerLocal(new(ctl.MapTestImpl), nil))
	mux.Handle("/pb", new(intf.PBH).HandlerLocal(new(ctl.PBImpl), nil))

	mux.Handle("/CaptchaID",
		new(intf.CaptchaIDH).HandlerLocal(new(ctl.CaptchaIDImpl), nil))
	mux.Handle("/CaptchaVerify",
		new(intf.CaptchaVerifyH).HandlerLocal(new(ctl.CaptchaVerifyImpl), nil))
	mux.Handle("/Captcha/",
		ycaptcha.Handler())

	//-------register micro-impl-----------------
	// 每二步:註冊微服務到etcd
	ylog.Info("正在向etcd注册微服务......")
	if err := yetcd.RegisterService(msTag, regHost, fmt.Sprint(regPort)); err != nil {
		ylog.Error("main.go->", err)
		return
	}
	ylog.Info("注册微服务", msTag, " 成功")

	//--------start server -------------------------
	//用于显示服务器状态，用于测试
	fmt.Println(fmt.Sprint("监听:", host, ":", port))
	testStr := fmt.Sprintf(
		`测试：curl -X POST  -H 'Content-Type:application/json'  -d '{"S":"from weihaoran /"}' %s:%d/CtlTest`,
		"127.0.0.1", port)
	fmt.Println(testStr)
	ylog.Info(fmt.Sprint("监听:", host, ":", port))

	addr := fmt.Sprint(host, ":", port)

	//-------------------------------------
	//  第三步：启动微服务
	golog.Fatal(http.ListenAndServe(addr, mux))
}

type Ping struct{}

func (r *Ping) handler(out http.ResponseWriter, in *http.Request) {
	_, _ = out.Write([]byte("pong........................"))
}
