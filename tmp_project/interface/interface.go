package main

import "fmt"

//https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/
//go 允许不带任何方法的 interface ，这种类型的 interface 叫 empty interface。

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s *S) Get() int {
	return s.Age
}
func (s *S) Set(age int) {
	s.Age = age
}

//i的参数类型是接口I
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	//S实现了I的两个方法，是I的实现者
	s := S{}
	f(&s)
}
