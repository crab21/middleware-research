package topic

import "time"

func TopicStart() {
	go TopicReceiver()

	time.Sleep(1*time.Second)

	go TopicSend("*.com.sp.yu")
	go TopicSend("*.com.ppyss.yu")
	select {

	}
}
