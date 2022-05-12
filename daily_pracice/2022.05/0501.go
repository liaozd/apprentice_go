package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages) // 关闭通道，禁止再向通道发送消息。这之后无法再向通道发送信息，会引发panic
	fmt.Println("Pushed two messages onto Channel with no receiver")
	time.Sleep(time.Second * 1)
	receiver(messages)
}
