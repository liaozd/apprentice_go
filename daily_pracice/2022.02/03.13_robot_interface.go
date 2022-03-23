package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

//Boot 接口也是一种类型，可作为参数传递给函数，因此可编写可重用于多个接口实现的函数。
//这个函数将接口Robot的实现作为参数，并返回调用方法PowerOn的结果。
//这个函数可用于启动任何机器人，而不管其方法PowerOn是如何实现的。T850和R2D2都可利用这个方法。
func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}

	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
