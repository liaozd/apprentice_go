package main

import (
	"fmt"
)

func main() {
	Young(17)
	Young(19)

	IfUsingNewVar(3, 105)
	IfUsingNewVar(3, 99)
}

func Young(age int) {
	if age < 18 {
		println(age, "JJ")
	} else {
		// else 可没有
		fmt.Println("I not a child")
	}

}

func IfUsingNewVar(start int, end int) {
	if distance := end - start; distance > 100 {
		fmt.Printf("距离大于100: %d\n", distance)
	} else {
		fmt.Printf("距离小于100: %d\n", distance)
	}

	// if else之外不可访问distance
	//fmt.Printf("距离是: %d\n", distance)

}
