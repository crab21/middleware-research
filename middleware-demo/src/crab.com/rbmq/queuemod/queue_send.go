package queuemod

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

func StartSendRBMQQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declare, err := ch.QueueDeclare("hello-work-test-queue", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	i:=0
	for {
		bodys:=body+strconv.Itoa(i)
		ch.Publish("", declare.Name, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(bodys),
		})
		i++
		time.Sleep(700 * time.Millisecond)
	}
	failOnError(err, "Failed to publish a message")

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
