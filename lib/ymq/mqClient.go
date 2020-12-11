package ymq

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"

	"github.com/vhaoran/vchat/lib/yconfig"
)

type Mqtt struct {
	URL      string
	token    MQTT.Token
	client   MQTT.Client
	clientID string

	Username string
	Password string
}

func NewMqtt(url string, userName, password string) (*Mqtt, error) {
	//??? extend
	var m = &Mqtt{
		URL:      url,
		Username: userName,
		Password: password,
	}

	err := m.connect()
	if err != nil {
		return nil, err
	}
	return m, nil
}

//Connect opens connection
func (m *Mqtt) connect() error {
	//uuid
	//uid, err := uuid.NewV1()
	uid := uuid.NewV1()
	////if err != nil {
	//	panic(err)
	//}
	//

	m.clientID = uid.String()

	opts := MQTT.NewClientOptions().AddBroker(m.URL)
	opts.SetClientID(m.clientID)

	opts.SetMaxReconnectInterval(time.Minute * 10)

	opts.Username = m.Username
	opts.Password = m.Password

	//opts.SetDefaultPublishHandler(m.messageHandler)
	m.client = MQTT.NewClient(opts)

	m.token = m.client.Connect()
	m.token.Wait()

	return m.token.Error()
}

//subscribe subscribes to topic
func (m *Mqtt) Subscribe(topic string, callback MQTT.MessageHandler) error {
	m.token = m.client.Subscribe(topic, 2, callback)
	m.token.Wait()
	return m.token.Error()
}

func (m *Mqtt) SubscribeQos(topic string, qos byte, callback MQTT.MessageHandler) error {
	m.token = m.client.Subscribe(topic, qos, callback)
	m.token.Wait()
	return m.token.Error()
}

//Publish publishes message to broker
func (m *Mqtt) Publish(topic string, msg interface{}) error {
	m.token = m.client.Publish(topic, 2, false, msg)
	m.token.Wait()
	return m.token.Error()
}

func (m *Mqtt) PublishQos(topic string, qos byte, msg interface{}) error {
	m.token = m.client.Publish(topic, qos, false, msg)
	m.token.Wait()
	return m.token.Error()
}

func (m *Mqtt) Destroy() error {
	//m.token = m.client.Unsubscribe("go-mqtt/sample")
	//m.token.Wait()
	m.client.Disconnect(250)
	return m.token.Error()
}

func getClient(cfg yconfig.MQConfig) (*Mqtt, error) {
	//if cfg == nil {
	//	return nil, errors.New("没有mq配置文件")
	//}

	url := cfg.Url
	if len(url) == 0 {
		url = fmt.Sprintf("%s:%s", cfg.Host, cfg.TCPPort)
	}

	return NewMqtt(url, cfg.UserName, cfg.Password)
}
