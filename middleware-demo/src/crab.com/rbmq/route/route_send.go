package route

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

func RouteSend() {
	conn, err := amqp.Dial("amqp://guest:guest@47.106.231.162:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs-direct", "direct", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")

	a:=1
	for   {
		sp:="aaaaaa"+strconv.Itoa(a)
		err = ch.Publish("logs-direct", "black", false, false, amqp.Publishing{Body: []byte(sp),
			ContentType: "text/plain"})
		failOnError(err, "Failed to publish a message")

		time.Sleep(100*time.Millisecond)
		a++
	}

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
