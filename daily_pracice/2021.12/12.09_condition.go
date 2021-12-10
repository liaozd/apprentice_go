package main

// Effective GO: control structures
import "fmt"

func main() {
	// if后面的大括号，不能换行，否则出错
	i := 1
	if i < 2 {
		fmt.Println("aaa")
	}

	fmt.Println("如果你只需要range中的第一项（key或者index），则可以丢弃第二个")
	array := []string{"a", "b", "c"}
	for index := range array {
		fmt.Printf("Keep only first in range: %d\n", index)
	}

	fmt.Println("错误的编码将占用一个字节，并以符文 U+FFFD 来代替。")
	for pos, char := range "日本\\x80語" {
		fmt.Printf("character %#U, char %d \n", char, pos)
	}

	fmt.Println("++, --, Go 没有逗号操作符，而 ++ 和 -- 为语句而非表达式。 因此，若你想要在 for 中使用多个变量，" +
		"应采用平行赋值的方式 （因为它会拒绝 ++ 和 --）.")
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}
