package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

type Bird struct{}

func (bird Bird) Fly() {
	fmt.Println("The bird is flying")
}

func (bird Bird) Run() {
	fmt.Println("The bird is running")
}

func main() {
	var animal Animal

	bird := new(Bird)

	// 将对象实例赋值给接口
	animal = bird
	animal.Fly()
	animal.Run()
}
