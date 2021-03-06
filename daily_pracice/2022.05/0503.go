package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 2)
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
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
	case <-time.After(500 * time.Microsecond):
		//通过使用超时时间，可在指定时间过后从select语句返回，从而结束阻塞操作。select语句根据最先到达的消息执行相应的case语句；
		//通过指定超时时间，可在给定时间内没有收到任何消息时从select语句返回。
		fmt.Println("超时500毫秒后执行")
	}
}
