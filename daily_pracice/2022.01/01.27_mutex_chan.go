package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("Already Locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Before lock: ", i)

			mutex.Lock()
			fmt.Println("Locked: ", i)

			time.Sleep(time.Second / 2)

			mutex.Unlock()
			fmt.Println("Unlock the lock: ", i)
			c <- i
		}(i, channels[i])
	}

	time.Sleep(time.Second / 2)
	mutex.Unlock()
	fmt.Println("Unlock the lock outside")
	time.Sleep(time.Second / 2)

	for _, c := range channels {
		fmt.Println("channel: ", c)
		<-c
	}
}
