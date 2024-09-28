package pingPong

import "time"

func ping(messages *chan string, seconds int) {
	for i := 1; i <= seconds; i++ {
		*messages <- "ping"
		time.Sleep(time.Second)
		message := <-*messages
		println(message)
	}
}

func pong(messages *chan string, seconds int) {
	for i := 1; i <= seconds; i++ {
		message := <-*messages
		println(message)
		time.Sleep(time.Second)
		*messages <- "pong"
	}
}

func PingPong(messages *chan string, seconds int) {
	go ping(messages, seconds)
	go pong(messages, seconds)

	time.Sleep(time.Duration(seconds+1) * time.Second)
}
