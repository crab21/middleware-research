package route

import "time"

func RouteStart() {
	go RouteReceiver()
	time.Sleep(1*time.Second)


	go RouteSend()
	go RouteSend()

	select {

	}
}
