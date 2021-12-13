package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 7
	i := x.(int)
	fmt.Println(reflect.TypeOf(i))
	value, ok := x.(int32)
	if !ok {
		fmt.Println("Type not ok")
		return
	}
	fmt.Println(reflect.TypeOf(value))
	// panic: interface conversion: interface {} is int, not int32
}
