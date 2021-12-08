package main

import (
	"fmt"
	"unsafe"
)

type Num struct {
	i string
	j int64
}

func main() {
	unsafeP()

	offSetOf()
}

func unsafeP() {
	num := 5
	numPointer := &num

	// 下面会报错: cannot convert numPointer (type *int) to type *float32
	//float_num := (*float32)(numPointer)
	//fmt.Println(float_num)

	float_num := (*float32)(unsafe.Pointer(numPointer))
	fmt.Println(float_num)
	// 结果: 0xc0000b2008
	fmt.Println(numPointer)
}

func offSetOf() {
	n := Num{i: "EDDDDDLK", j: 1}
	nPointer := unsafe.Pointer(&n)

	// - 结构体的成员变量在内存存储上是一段连续的内存
	// - 结构体的初始地址就是第一个成员变量的内存地址
	// - 基于结构体的成员地址去计算偏移量。就能够得出其他成员变量的内存地址

	// 修改 n.i 值：i 为第一个成员变量。因此不需要进行偏移量计算，直接取出指针后转换为 Pointer，再强制转换为字符串类型的指针值即可
	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "煎鱼"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 2

	fmt.Printf("n.i: %s, n.j: %d \n", n.i, n.j)

}
