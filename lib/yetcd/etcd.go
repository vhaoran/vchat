package yetcd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"

	"github.com/vhaoran/vchat/lib/yconfig"
	"github.com/vhaoran/vchat/lib/ylog"
)

/*--auth: whr  date:2019/12/511:45--------------------------
 ####请勿擅改此功能代码####
 用途：
 用于服务自动发现
 --->yconfig
--------------------------------------- */

var (
	XETCDConfig = &yconfig.ETCDConfig{
		Hosts: []string{"127.0.0.1:2379"},
		Options: etcdv3.ClientOptions{
			Cert:          "",
			Key:           "",
			CACert:        "",
			DialTimeout:   time.Second * 3,
			DialKeepAlive: time.Second * 3,
			Username:      "",
			Password:      "",
		},
	}
)

//初始化 ETCD 配置
//todo
func InitETCD(cfg yconfig.ETCDConfig) error {
	//get config and fill to XETCDConfig
	*XETCDConfig = cfg
	//
	XETCDConfig.Options.DialTimeout = time.Second * 10
	XETCDConfig.Options.DialKeepAlive = time.Second * 100
	return nil
}

func RegisterService(serviceName, host, port string) error {
	ylog.Debug("--------etcd.go--->微服务注册中---", serviceName, ",", host, ",", port)

	ctx := context.Background()
	//options := etcdv3.ClientOptions{
	//	DialTimeout:   time.Second * 3,
	//	DialKeepAlive: time.Second * 3,
	//}
	options := etcdv3.ClientOptions{
		Cert:          XETCDConfig.Options.Cert,
		Key:           XETCDConfig.Options.Key,
		CACert:        XETCDConfig.Options.CACert,
		DialTimeout:   XETCDConfig.Options.DialTimeout * time.Second,
		DialKeepAlive: XETCDConfig.Options.DialKeepAlive * time.Second,
		Username:      XETCDConfig.Options.Username,
		Password:      XETCDConfig.Options.Password,
	}
	//set time to second
	//options.DialTimeout = options.DialTimeout * time.Second
	//options.DialKeepAlive = options.DialKeepAlive * time.Second

	etcdHosts := XETCDConfig.Hosts

	addr := fmt.Sprint(host, ":", port)
	client, err := etcdv3.NewClient(ctx, etcdHosts, options)
	if err != nil {
		ylog.Error("etcd->", err)
		return err
	}

	service := etcdv3.Service{
		Key:   fmt.Sprint(serviceName, "/", addr),
		Value: addr,
		TTL:   etcdv3.NewTTLOption(time.Second*3, time.Second*10),
	}
	if err = client.Register(service); err != nil {
		log.Println("###register micro unit error:", err)
		return err
	}

	log.Println(service.Key, "--register micro unit  successful....")
	return nil
}

func GetEtcdClient() (etcdv3.Client, error) {
	ctx := context.Background()
	options := etcdv3.ClientOptions{
		Cert:          XETCDConfig.Options.Cert,
		Key:           XETCDConfig.Options.Key,
		CACert:        XETCDConfig.Options.CACert,
		DialTimeout:   XETCDConfig.Options.DialTimeout * time.Second,
		DialKeepAlive: XETCDConfig.Options.DialKeepAlive * time.Second,
		Username:      XETCDConfig.Options.Username,
		Password:      XETCDConfig.Options.Password,
	}

	etcdHosts := XETCDConfig.Hosts

	//addr := fmt.Sprint(host, ":", port)
	client, err := etcdv3.NewClient(ctx, etcdHosts, options)

	return client, err
}
