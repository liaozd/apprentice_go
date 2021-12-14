package main

import "fmt"

func main() {
	type Transform [3][3]float64 // 一个 3x3 的数组，其实是包含多个数组的一个数组。
	type LinesOfText [][]byte    // 包含多个字节切片的一个切片。

	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}

	fmt.Println(text)
}
