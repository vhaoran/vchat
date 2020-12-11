package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/atomic"

	"github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ymqa"
)

func init() {
	_, err := lib.InitModulesOfOptions(&lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadRabbitMq:     true,
		LoadJwt:          false,
		LoadQiniu:  true,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Test_publish(t *testing.T) {
	for i := 0; i < 100; i++ {
		topic := fmt.Sprint("t_", i%5)
		s := fmt.Sprint("  ", i, " hello_", ytime.OfNow().String())
		//
		if err := ymqa.X.Publish(topic, s); err != nil {
			ylog.Error("amq_test.go->", err)
		} else {
			ylog.Debug(i, "  ok.....")
		}
	}
	time.Sleep(time.Second * 20)
}

func Test_Consume_1(t *testing.T) {
	f := func(body amqp.Delivery) error {
		defer func() {
			if err := recover(); err != nil {
				ylog.Error("call back error: ", err)
			}
		}()
		fmt.Println("---####--", time.Now(), " ", body.RoutingKey, "----####---", string(body.Body))
		return nil
	}

	c := make(chan int)
	for i := 0; i < 5; i++ {
		topic := fmt.Sprint("t_", i%5)
		if _, err := ymqa.X.Consume(topic, f, 10); err != nil {
			ylog.Debug(err)
		} else {
			ylog.Debug(i, "   consume ok  ", i)
		}
	}
	<-c
}

func Test_Consume_ack(t *testing.T) {
	k := atomic.NewInt32(0)

	f := func(obj amqp.Delivery) error {
		defer func() {
			if err := recover(); err != nil {
				ylog.Error("call back error: ", err)
			}
		}()
		fmt.Println("---####--", time.Now(), " ", obj.RoutingKey, "----####---", string(obj.Body))

		k.Inc()
		if k.Load()%2 == 1 {
			err := obj.Ack(false)
			if err != nil {
				ylog.Error("amq_test.go->ack err:", err)
				return err
			}
		}
		return nil
	}

	c := make(chan int)
	for i := 0; i < 5; i++ {
		topic := fmt.Sprint("t_", i%5)
		if _, err := ymqa.X.ConsumeAck(topic, f, 10); err != nil {
			ylog.Debug(err)
		} else {
			// ylog.Debug(i, "   consume ok  ", i)
		}
	}
	log.Println("............waiting......... .....")
	<-c
}

func Test_bench_publish(t *testing.T) {
	for i := 0; i < 100000; i++ {
		topic := fmt.Sprint("t_", i%5)
		s := fmt.Sprint("  ", i, " hello_", ytime.OfNow().String())
		//
		if err := ymqa.X.Publish(topic, s); err != nil {
			ylog.Error(err)
			ylog.Error("amq_test.go->", err)
		} else {
			ylog.Debug(i, "  ok.....")
		}
	}

	time.Sleep(time.Second * 20)
}
