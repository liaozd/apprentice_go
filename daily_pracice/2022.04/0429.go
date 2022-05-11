package main

import (
	"fmt"
	"time"
)

func slowFunc2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "slowFunc() Finished"
}

func main() {
	c := make(chan string)
	go slowFunc2(c)
	fmt.Println("after go slowFunc2")
	msg := <-c
	fmt.Println(msg)
}
