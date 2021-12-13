package main

import "fmt"

func main() {
	first := "first"
	fmt.Println([]rune(first))
	fmt.Println([]byte(first))

	first = "社区"
	fmt.Println([]rune(first))
	fmt.Println([]byte(first))
}
