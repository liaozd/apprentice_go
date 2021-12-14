package main

import "fmt"

// func append(slice []T, elements ...T) []T
//
func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
