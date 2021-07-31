package main

import "fmt"

func filter(score []int, f func(int) bool) []int {
	reSlice := make([]int, 0)

	for _, v := range score {
		if f(v) {
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}

func main() {
	myFunc := func(a, b int) int {
		return a + b
	}
	myFunc(1, 2)
	fmt.Printf("%T", myFunc)
	score := []int{10, 50, 80, 90, 85}
	fmt.Println(filter(score, func(a int) bool {
		if a >= 90 {
			return true
		} else {
			return false
		}
	}))

}
