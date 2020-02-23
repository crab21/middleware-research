package pubsub

import "time"

func PubsubStart() {
	go  PubSubReceiver()
	time.Sleep(1*time.Second)

	go PubsubSend()
	go PubsubSend()
	go PubsubSend()

	select {
	}
}
