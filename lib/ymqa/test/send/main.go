package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://root:password@192.168.0.99:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	topic := fmt.Sprint("t_", 3%5)

	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	failOnError(err, "Failed to declare a queue")

	t0 := time.Now()
	for k := 0; k < 10; k++ {
		i := k

		body := fmt.Sprint("Hello World!_", i, "---")
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
		time.Sleep(100 * time.Millisecond)

	}
	log.Println("time: ", time.Since(t0))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
