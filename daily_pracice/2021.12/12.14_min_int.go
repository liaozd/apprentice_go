package main

import "fmt"

func main() {
	fmt.Println(Min(1, 11, 111, 11111))
}

// 顺便一提，... 形参可指定具体的类型，例如从整数列表中选出最小值的函数 min，其形参可为 ...int 类型。
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // 最大的 int: 9223372036854775807
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
