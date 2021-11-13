package main

import (
	"fmt"
	_ "fmt"
) // 匿名引用

var a byte = 12

func main() {
	println("Hello, GO!")
	println(a)

	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)
	IfUsingNewVariable(10, 40)
}

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("Too Far: %d\n", distance)
	} else if distance < 50 {
		fmt.Printf("Too close: %d\n", distance)
	}
}
