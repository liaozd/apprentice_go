package main

import "fmt"

func main() {
	ForLoop()
	ForI()
	ForR()
}

func ForLoop() {
	arr := []int{9, 8, 7, 6}
	index := 0
	for {
		if index == 3 {
			break
		}
		fmt.Printf("%d => %d \n", index, arr[index])
		index++
	}
	fmt.Println("for loop end \n")
}

func ForI() {
	arr := []int{9, 8, 7, 6}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d => %d \n", i, arr[i])
	}
	fmt.Println("for i loop end \n ")
}

func ForR() {
	arr := []int{9, 8, 7, 6}

	for index, value := range arr {
		fmt.Printf("%d => %d \n", index, value)
	}

	for _, value := range arr {
		fmt.Printf("only value: %d \n", value)
	}

	for index := range arr {
		fmt.Printf("only index %d \n", index)
	}

	fmt.Printf("for r loop end \n ")
}
