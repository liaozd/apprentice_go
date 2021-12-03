package main

import "fmt"

func main() {
	// 初始化三个元素数组
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1: %v, len: %d, cap: %d\n", a1, len(a1), cap(a1))
	// 初始化三个0元素数组
	var a2 [3]int
	fmt.Printf("a2: %v, len: %d, cap: %d\n", a2, len(a2), cap(a2))

	//使用索引
	fmt.Printf("a1[1]: %d", a1[1])

	//超出索引会编译出错
	//fmt.Printf("a1[99]: %d", a1[99])
}
