package main

import (
	"fmt"
	"time"
)

// 我们在Goroutine中创建一个函数，它每隔1s向通道发送一条消息。无限不终止
func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C // C <-chan Time | The channel on which the ticks are delivered.
	}
}

func main() {
	messages := make(chan string)
	stop := make(chan bool)

	go sender(messages)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Time's up!")
		stop <- true
	}()

	// 通过在for循环中使用select语句，可在收到消息后立即打印它。由于这是一个阻塞操作，因此将不断打印消息，直到您手动终止这个过程。
	for {
		select {
		case <-stop: // 接收 stop channels的消息
			return
		case msg := <-messages: // 接收通道每秒一次的信息
			fmt.Println(msg)
		}
	}
}
