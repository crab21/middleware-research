package workmod

import "time"

func WorkStart() {
	go StartReceiver()
	time.Sleep(1*time.Second)

	StartSendRBMQ()
	select {

	}
}
