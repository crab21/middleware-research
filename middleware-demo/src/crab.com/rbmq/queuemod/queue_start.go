package queuemod

import "time"

func StartQueue() {
	go StartReceiverQueue()
	time.Sleep(1*time.Second)

	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	go StartSendRBMQQueue()
	select {

	}
}