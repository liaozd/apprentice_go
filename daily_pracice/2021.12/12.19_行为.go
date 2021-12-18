package main

import (
	"fmt"
	"unsafe"
)

type Employee2 struct {
	Id   string
	Name string
	Age  int
}

func (e Employee2) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee2) String2() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func main() {
	e := Employee2{"0", "Bob", 20}
	// 实例可调用方法
	fmt.Printf("e.String(): %v\n", e.String())

	// 实例指针也可调方法
	eP := &Employee2{"0", "Bob", 20}
	fmt.Printf("Pointer e.String(): %v\n", eP.String())
}
