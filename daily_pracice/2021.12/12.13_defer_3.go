package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
