package main

import "Desafio-Dio-Ping-Pong-Golang/pingPong"

func main() {

	messages := make(chan string)

	pingPong.PingPong(&messages, 5)
}
