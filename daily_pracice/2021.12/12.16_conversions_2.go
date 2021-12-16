package main

import (
	"fmt"
	"sort"
)

type SequenceSimple []int

// 用于打印的方法 - 在打印前对元素进行排序。
func (s SequenceSimple) String() string {
	// 不必让 Sequence 实现多个接口（排序和打印）
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main() {
	s := SequenceSimple{1, 2, 3}
	fmt.Println(s)
}
