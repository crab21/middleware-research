package workmod

import (
	"github.com/streadway/amqp"
	"log"
)

func StartReceiver() {

	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()


	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declare, err := ch.QueueDeclare("hello-work-test", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")


	consume, err := ch.Consume(declare.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d:=range consume {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


}