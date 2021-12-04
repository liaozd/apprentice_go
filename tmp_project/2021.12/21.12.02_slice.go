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
	fmt.Printf("s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	// 按索引取值
	fmt.Printf("s3[2] %d \n", s3[2])

	// 超出下标范围，溢出
	//fmt.Printf("s3[99]", s3[99])

	SubSlice()

	ShareSlice()

}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}

	// 长度2, cap 4
	s2 := s1[1:3]
	fmt.Printf("subslice s2: %v, len %d, cap: %d \n", s2, len(s2), cap(s2))

	// 长度3, cap 4
	s3 := s1[2:]
	fmt.Printf("subslice s3: %v, len %d, cap: %d \n", s3, len(s3), cap(s3))

	// 长度3, cap 5
	s4 := s1[:3]
	fmt.Printf("subslice s4: %v, len %d, cap: %d \n", s4, len(s4), cap(s4))
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1: %v, len %d, cap: %d address: %p \n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2: %v, len %d, cap: %d address: %p \n", s2, len(s2), cap(s2), &s2)
	//s1: [1 2 3 4], len 4, cap: 4 address: 0xc00000c108
	//s2: [3 4], len 2, cap: 2 address: 0xc00000c120

	s2[0] = 99
	fmt.Printf("s1: %v, len %d, cap: %d address: %p \n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2: %v, len %d, cap: %d address: %p \n", s2, len(s2), cap(s2), &s2)
	//s1: [1 2 99 4], len 4, cap: 4 address: 0xc00000c108
	//s2: [99 4], len 2, cap: 2 address: 0xc00000c120

	s2 = append(s2, 199)
	fmt.Printf("s1: %v, len %d, cap: %d address: %p \n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2: %v, len %d, cap: %d address: %p \n", s2, len(s2), cap(s2), &s2)
	//s1: [1 2 99 4], len 4, cap: 4 address: 0xc00000c108
	//s2: [99 4 199], len 3, cap: 4 address: 0xc00000c120

	s2[1] = 1999
	fmt.Printf("s1: %v, len %d, cap: %d address: %p \n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("s2: %v, len %d, cap: %d address: %p \n", s2, len(s2), cap(s2), &s2)
	//s1: [1 2 99 4], len 4, cap: 4 address: 0xc00000c108
	//s2: [99 1999 199], len 3, cap: 4 address: 0xc00000c120  <-- 这块不知道为什么, s2[1]为什么没有替换掉s1[2]
}
