package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = 7
	i := x.(int)
	fmt.Println(reflect.TypeOf(i))
	j := x.(int32)
	fmt.Println(reflect.TypeOf(j))
	// panic: interface conversion: interface {} is int, not int32
}
