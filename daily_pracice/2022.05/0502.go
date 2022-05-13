package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C // 用于实现每秒钟发送一条
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)
	for i := 0; i < 5; i++ {
		msg := <-messages
		fmt.Println(msg)
	}
}
