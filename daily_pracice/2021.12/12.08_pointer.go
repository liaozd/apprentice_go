package main

import "fmt"

func main() {
	var p *ToyDuck = &ToyDuck{}
	var duck ToyDuck = *p
	duck.Swim()

	var nilDuck *ToyDuck
	if nilDuck == nil {
		fmt.Println("nikDuck is nil")
	}
}
