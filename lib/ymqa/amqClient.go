package ymqa

import (
	"errors"
	"time"

	"github.com/streadway/amqp"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/lib/ylog"
)

func getCntOfRabbitMQ(url string) (*amqp.Connection, error) {
	//conn, err := amqp.Dial("amqp://root:password@localhost:5672/")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func publishWrap(conn *amqp.Connection, data *AMqData) error {
	ch, err := conn.Channel()
	if err != nil {
		return errors.New("没有获取到rabbit连接")
	}
	defer func() {
		_ = ch.Close()
	}()

	queue := data.Queue
	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		ylog.Error("amqClient.go->publishWrap,err:", err)
		return err
	}

	body, er1 := g.GetBufferForMq(data.Data)
	if er1 != nil {
		ylog.Error("amqClient.go->publishWrap", er1)
		return er1
	}
	if len(body) == 0 {
		ylog.Error("amqClient.go->unmarshal数据为空")
		return er1
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	return err
}

func consumeWrap(conn *amqp.Connection, queue string, callback AMQSubCallBack, autoAck bool) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	l, er1 := ch.Consume(
		q.Name,  // queue
		"",      // consumer
		autoAck, // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if er1 != nil {
		ylog.Error("amqClient.go->consumeWrap", er1)
		return er1
	}

	//forever := make(chan bool)
	go func() {
		for {
			select {
			case d, ok := <-l:
				if !ok {
					time.Sleep(50 * time.Millisecond)
					continue
				}
				func() {
					defer func() {
						if err := recover(); err != nil {
							ylog.Error("amqClient.go->consumeWrap->deferCallback", err)

						}
					}()
					_ = callback(d)
				}()
			}
		}
	}()

	return nil
}
