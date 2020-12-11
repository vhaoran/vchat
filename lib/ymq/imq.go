package ymq

import (
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X *MqWorker
)

type (
	IMq interface {
		Publish(nodeID int64, topic string, msg interface{}) error
		PublishQos(nodeID int64, topic string, qos byte, msg interface{}) error
		Subscribe(nodeID int64, topic string, callback MQTT.MessageHandler) (cnt *Mqtt, err error)
	}
)

func InitMq(cfg yconfig.MQConfig) error {
	cnt := new(MqWorker)
	err := cnt.NewPoolClient(&cfg)
	if err != nil {
		log.Println(err)
		return err
	}

	X = cnt
	return nil
}
