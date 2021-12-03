package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4} // 初始化4个元素的切片
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	var s12 = []int{}
	fmt.Printf("s12: %v, len: %d, cap: %d \n", s12, len(s12), cap(s12))

	s2 := make([]int, 3, 4) // 创建容量为4的切片
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// s2
	s2 = append(s2, 7) // s2后面追加一个元素，cap不变
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// 添加一个元素，触发扩容
	s2 = append(s2, 8)
	fmt.Printf("s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	s3 := make([]int, 4) // 只传入一个参数，表示创建4元素
	fmt.Printf("s3: %v", s3)

}
