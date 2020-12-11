package ymq

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/atomic"

	//"chat/Library/Inner/Sources/yconfig"
	//"chat/Library/Inner/Sources/ylog"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib/yconfig"
)

type (
	MqWorker struct {
		sync.RWMutex
		//当前已连接数量
		count atomic.Int64
		//最小数量
		min atomic.Int64

		//最大数量
		max atomic.Int64
		//chan,消息传递队列
		queue chan *MqData

		config *yconfig.MQConfig
	}

	MqData struct {
		Topic string
		//0/1/2
		Qos byte
		//must be bytes or string
		Data interface{}
	}
)

const (
	chanLength = 2000000
)

func (r *MqWorker) NewPoolClient(config *yconfig.MQConfig) error {
	r.Lock()
	r.config = config
	r.queue = make(chan *MqData, chanLength)
	r.Unlock()

	r.min.Store(int64(config.MinOpenConnS))
	r.max.Store(int64(config.MaxOpenConnS))
	r.count.Store(0)

	go r.scan(*config)
	return r.createOne(*config)
}

func (r *MqWorker) scan(cfg yconfig.MQConfig) {
	for {
		count := r.count.Load()
		min := r.min.Load()
		max := r.max.Load()

		//log.Println("##### emq pool count:", count, " max:", max, " lenOfQueue(chan)", lenOfQueue(r.queue))
		if count < min {
			_ = r.createOne(cfg)
			continue
		}

		//如果长度过半，则加一个长度
		lenOfQueue := len(r.queue)
		if int64(lenOfQueue) > count*5 && count < max {
			//ylog.SysLogN().WarnF("emq连接池数量需要增加 count: %d lenOfQueue: %d  max: %d", count, lenOfQueue, max) //log.Printf("emq连接池数量需要增加 count: %d lenOfQueue: %d, max: %d", count, lenOfQueue, max)
			_ = r.createOne(cfg)
			continue
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func (r *MqWorker) createOne(cfg yconfig.MQConfig) error {
	cnt, err := r.GetCntDirect(cfg)
	//
	if err != nil {
		return errors.New("没有获取到emq连接")
	}
	r.count.Inc()
	log.Println("mq pool added,len (", r.count.Load())

	go func() {
		//出错时，释放连接数量
		defer func() {
			r.count.Dec()
			_ = cnt.Destroy()
			log.Println("##### mq pool close,len(", r.count.Load(), ")")
		}()

		//start := time.Now()
		endLoop := false
		for !endLoop {
			select {
			case bean, ok := <-r.queue:
				if ok && bean != nil {
					success := false
					for i := 0; i < 3; i++ {
						if err := cnt.PublishQos(bean.Topic, bean.Qos, bean.Data); err == nil {
							success = true
							break
						}
						time.Sleep(50 * time.Millisecond)
					}

					// 重新放入对列中，进行发送
					if !success {
						_ = r.PublishQos(bean.Topic, bean.Qos, bean.Data)
					}
					continue
				}
			case <-time.After(time.Minute * 10):
				log.Println("mq wait timeout-->,and recycle...")
				if r.count.Load() > r.min.Load() {
					endLoop = true
				}
			}
		}
	}()
	return nil
}

func (r *MqWorker) getSleepDuration(connCount int64) time.Duration {
	if connCount <= 50 {
		return time.Millisecond * 500
	}

	if connCount <= 100 {
		return time.Second * 1000
	}
	if connCount > 1000 {
		return time.Second * 10000
	}
	return time.Duration(connCount/100) * time.Second
}

func (r *MqWorker) Publish(topic string, msg interface{}) error {
	return r.PublishQos(topic, 1, msg)
}

func (r *MqWorker) PublishQos(topic string, qos byte, msg interface{}) error {
	buffer, err := g.GetBufferForMq(msg)
	if err != nil {
		return err
	}

	bean := &MqData{
		Topic: topic,
		Qos:   qos,
		Data:  string(buffer),
	}

	select {
	case r.queue <- bean:
	case <-time.After(100 * time.Millisecond):
		return errors.New("MqWorker) PublishQos发送消息超时")
	}
	return nil
}

func (r *MqWorker) Subscribe(topic string, handler mqtt.MessageHandler) (cnt *Mqtt, err error) {
	//---------从池中扑出来一个，不归还的mq client------------
	cnt, err = r.GetCntDirect(*r.config)
	if cnt == nil {
		return nil, fmt.Errorf("emqx Subscribe 无法再建立连接")
	}

	//订阅并返回，有多个订阅时，可以直接在返回结果上操作
	err = cnt.SubscribeQos(topic, 2, handler)
	return cnt, err
}

// 根据配置，直接生成一个新的连接
func (r *MqWorker) GetCntDirect(cfg yconfig.MQConfig) (cnt *Mqtt, err error) {
	// lock ??
	cnt, err = getClient(cfg)
	return
}
