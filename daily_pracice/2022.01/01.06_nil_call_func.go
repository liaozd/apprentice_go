package main

import "fmt"

type T struct{}

func (t *T) Hello() string {
	fmt.Printf("%v, %T, %x\n", t, t, &t)
	// 输出: <nil>, *main.T, c0000ac028
	if t == nil {
		fmt.Println("NIL CALLED")
		return ""
	}

	return "哈哈"
}
func main() {
	var t *T

	fmt.Printf("%v, %T, %x\n", t, t, &t)
	// 输出: <nil>, *main.T, c0000ac018

	// 表达式 Expression.Name 的语法，所调用的函数完全由 Expression 的类型决定。
	// 其调用函数的指向不是由该表达式的特定运行时值来决定，包括我们前面所提到的 nil。
	// func (p *Sometype) Somemethod (firstArg int) {} 本质上是 -> func SometypeSomemethod(p *Sometype, firstArg int) {}

	s := t.Hello()
	fmt.Println(s)
}
