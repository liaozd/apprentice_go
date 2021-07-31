package main

import "fmt"

func main() {
	var a [4]int
	a[0] = 1
	fmt.Println(a)

	b := [2]string{"Penn", "Teller"}
	c := []string{"Penn", "Teller"}
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
