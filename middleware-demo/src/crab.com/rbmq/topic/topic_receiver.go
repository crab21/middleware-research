package topic

import (
	"github.com/streadway/amqp"
	"log"
)

func TopicReceiver() {
	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs-topic", "topic", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare("", false, false, true, false, nil)

	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "#.yu", "logs-topic", false, nil)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)

		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

	select {}
}

