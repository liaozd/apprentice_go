package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("in go func()")
		messages <- "ping"
		fmt.Println("already sent")
	}()

	msg := <-messages
	fmt.Println("after chan msg")
	fmt.Println("got channel msg: ", msg)
}
