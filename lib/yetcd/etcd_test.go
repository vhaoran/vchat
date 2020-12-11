package yetcd

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/sd/etcdv3"
)

func init() {

}

func Test_register(t *testing.T) {
	ctx := context.Background()
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	client, err := etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, options)
	fmt.Println("-----------------", client, err)
	fmt.Println("------aaa-----------")
	service := etcdv3.Service{
		Key:   "/hello",
		Value: "http://127.0.0.1:8888/",
		TTL:   etcdv3.NewTTLOption(time.Second*3, time.Second*10),
	}
	err = client.Register(service)
	//
	service = etcdv3.Service{
		Key:   "/hello1",
		Value: "http://127.0.0.1:8889/",
		TTL:   etcdv3.NewTTLOption(time.Second*3, time.Second*10),
	}
	err = client.Register(service)

	if err != nil {
		log.Println("err:", err)
	} else {
		log.Println("ok")
	}

	lst, err1 := client.GetEntries("/hello")
	fmt.Println("--------values---------")
	spew.Dump("lst:", lst)
	spew.Dump("err:", err1)
	time.Sleep(time.Second * 5)
}

func Test_mult(t *testing.T) {
	ctx := context.Background()
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}

	l := make([]*etcdv3.Client, 0)

	var wg sync.WaitGroup
	wg.Add(100)
	client, err := etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, options)
	l = append(l, &client)
	//
	go func() {
		for i := 0; i < 100; i++ {
			client, err = etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, options)

			service := etcdv3.Service{
				Key:   fmt.Sprint("/hello", i),
				Value: fmt.Sprint("http://127.0.0.1:888--", i),
				TTL:   etcdv3.NewTTLOption(time.Second*3, time.Second*10),
			}
			if err = client.Register(service); err != nil {
				fmt.Println("###register error:", err)
			} else {
				fmt.Println(service.Key, "--register ok")
				fmt.Println("client.LeaseID():  ", client.LeaseID())
			}
			time.Sleep(time.Second * 5)
			wg.Done()
		}
	}()

	go watch_visit()

	wg.Wait()
}

//func RegisterServiceX(serviceName, host, port string) error {
//	ctx := context.Background()
//	options := etcdv3.ClientOptions{
//		DialTimeout:   time.Second * 3,
//		DialKeepAlive: time.Second * 3,
//	}
//
//	addr := fmt.Sprint(host, ":", port)
//	client, err := etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, options)
//	unit := etcdv3.Service{
//		Key:   fmt.Sprint(serviceName, "||", addr),
//		Value: addr,
//		TTL:   etcdv3.NewTTLOption(time.Second*3, time.Second*10),
//	}
//	if err = client.Register(unit); err != nil {
//		fmt.Println("###register error:", err)
//		return err
//	}
//
//	fmt.Println(unit.Key, "--register successful")
//	return nil
//}

func watch_visit() {
	fmt.Println("###################### enter visit#########")
	ctx := context.Background()
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 10,
		DialKeepAlive: time.Second * 50,
	}
	client, err := etcdv3.NewClient(ctx, []string{"127.0.0.1:2379"}, options)
	if err != nil {
		fmt.Println("err watch client", err)
		return
	}

	ch := make(chan struct{})
	go client.WatchPrefix("/hello", ch)
	for {
		select {
		case <-ch:
			if instances, err := client.GetEntries("/hello"); err == nil {
				for _, v := range instances {
					fmt.Println(v)
				}
				fmt.Println("### all count ", len(instances), " #####")
			} else {
				fmt.Println("watch loop err:", err)
			}
		}
	}
}
