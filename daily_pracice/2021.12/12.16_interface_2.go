package main

import "fmt"

type Animal1 interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

type Bird1 struct{}

func (bird Bird1) Fly() {
	fmt.Println("the bird is flying")
}

func (bird Bird1) Run() {
	fmt.Println("the bird is running")
}

func main() {
	var animal1 Animal1
	var animal2 Animal2
	bird1 := new(Bird1)

	animal1 = bird1

	// 将animal1接口赋值给animal2，注意函数功能包含关系，包含方法多的接口可以赋值给方法少的接口，反之，则不行。
	animal2 = animal1

	animal1.Fly()
	animal1.Run()
	animal2.Fly()

}
