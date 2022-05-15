package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping2(channel2)
	go ping1(channel1)

	//假设有多个Goroutine，而程序将根据最先返回的Goroutine执行相应的操作，此时可使用select语句。
	//select语句类似于第5章介绍的switch语句，它为通道创建一系列接收者，并执行最先收到消息的接收者。
	select {
	case msg1 := <-channel1:
		fmt.Println("received", msg1)
	case msg2 := <-channel1:
		fmt.Println("received", msg2)
	}
}
