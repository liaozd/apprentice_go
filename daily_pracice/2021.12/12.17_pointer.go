package main

import "fmt"

type Employee struct {
	Id string
	Name string
	Age int
}

func main() {
	e := Employee{"0", "P1", 33}
	eNewed := new(Employee)  // new返回的是一个指针
	eNewed.Id = "1"
	fmt.Printf("e: %T\n", e)
	// e: main.Employee 类型是一个
	fmt.Printf("e: %T\n", eNewed)
	// e: *main.Employee
}