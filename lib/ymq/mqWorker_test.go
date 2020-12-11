package ymq

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib/yconfig"
)

func Test_mq_publish(t *testing.T) {
	cfg := &yconfig.MQConfig{
		Url:          "",
		Host:         "192.168.0.201",
		TCPPort:      "1883",
		UserName:     "",
		Password:     "",
		MinOpenConnS: 200,
		MaxOpenConnS: 500,
	}

	cnt := new(MqWorker)
	err := cnt.NewPoolClient(cfg)
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(10)
	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		msg := fmt.Sprint(time.Now(), "---msg--", i)
		_ = cnt.PublishQos("test", 2, "hello "+msg)
	}
	log.Println("times:", time.Since(t0))
	//
	wg.Wait()
}
