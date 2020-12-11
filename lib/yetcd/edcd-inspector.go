package yetcd

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib/ylog"
)

func InspectLoop(serviceName, host, port string) {
	go loop(serviceName, host, port)
}

func loop(serviceName, host, port string) {
	addr := fmt.Sprint(host, ":", port)
	key := fmt.Sprint(serviceName, "/", addr)
	ylog.Debug("--------edcd-inspector.go--->enter")
	fmt.Println(key, addr)

	var c etcdv3.Client
	var err error

	//--------redo -----------------------------
REDO:
	err = RegisterService(serviceName, host, port)
	if err != nil {
		time.Sleep(time.Second * 1)
		ylog.Debug("--------edcd-inspector.go--->register error", err)
		goto REDO
	}

	//-------- -----------------------------
	c, err = GetEtcdClient()
	if err != nil {
		time.Sleep(time.Second * 1)
		ylog.Debug("--------edcd-inspector.go--->register error", err)
		goto REDO
	}

	//
	//go
	ylog.Debug("--------edcd-inspector.go--->before watch")

	ylog.Debug("--------edcd-inspector.go--->after watch")
	//--------loop -----------------------------
	for {
		time.Sleep(time.Second * 10)

		l, err := c.GetEntries(key)
		if err != nil {
			ylog.Debug("--------edcd-inspector.go--not get entries->---")
			goto REDO
		}

		if !g.InSlice(addr, l) {
			for _, v := range l {
				ylog.Debug("----******----edcd-inspector.go------", v)
			}

			goto REDO
		}

		for _, v := range l {
			ylog.Debug("--------etcd working ok---", v)
		}
	}
}
