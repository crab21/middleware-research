package workmod

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

func StartSendRBMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	declare, err := ch.QueueDeclare("hello-work-test", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	//for {
		ch.Publish("",declare.Name,false,false,amqp.Publishing{

			ContentType:     "text/plain",
			Body:            []byte(body),
		})
		time.Sleep(10*time.Millisecond)
	//}
	failOnError(err, "Failed to publish a message")
	select {

	}

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
