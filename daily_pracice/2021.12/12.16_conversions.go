package main

import (
	"fmt"
	"sort"
)

type SequenceCon []int

func (s SequenceCon) Len() int {
	return len(s)
}

func (s SequenceCon) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SequenceCon) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SequenceCon) String() string {
	sort.Sort(s)
	// 转换过程并不会创建新值，它只是值暂让现有的时看起来有个新类型而已。
	//（还有些合法转换则会创建新值，如从整数转换为浮点数等。）
	return fmt.Sprint([]int(s))
}
func main() {
	s := SequenceCon{1, 2, 3}
	fmt.Println(s)
}
