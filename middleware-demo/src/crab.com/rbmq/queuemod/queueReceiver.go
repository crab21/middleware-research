package queuemod

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func StartReceiverQueue() {

	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()


	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declare, err := ch.QueueDeclare("hello-work-test-queue", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(10, 0, false)
	failOnError(err, "Failed to set QoS")

	consume, err := ch.Consume(declare.Name, "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	location, err := time.LoadLocation("Asia/Shanghai")

	go func() {
		for d:=range consume {
			log.Println(time.Now().In(location).UnixNano())
			log.Printf("Received a message: %s", d.Body)
			//time.Sleep(1*time.Second)
			d.Ack(false)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


}
