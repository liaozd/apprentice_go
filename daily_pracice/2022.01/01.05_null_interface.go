package main

import "fmt"

func DoSth(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("integer", i)
		return
	}
	if i, ok := p.(string); ok {
		fmt.Println("string", i)
		return
	}
	fmt.Println("Unknown Type")
}
