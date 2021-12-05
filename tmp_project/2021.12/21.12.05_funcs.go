package main

import "fmt"

func main() {
	a := Fun0("Tom")
	println(a)

	b, c := Fun1("a", 12)
	println(b, c)

	age, name := Fun2("a string", "b string")
	println(age, name)

	var d, e, g = Fun3("a", "b", "c", 11, 11, "10")
	println(d, e, g)

	Fun4("hello", 19, "A", "D")
	s := []string{"a", "b", "c"}
	Fun4("hello", 19, s...)
}

// Fun0 单一返回值，不需要括号
func Fun0(name string) string {
	return "Hello, " + name
}

// Fun1 多参数、多返回值， 参数有名字，返回值没有
func Fun1(a string, b int) (int, string) {
	return 0, "你好"
}

// Fun2 的返回值标注名字，可以直接在内部直接复制，然后返回
func Fun2(a string, b string) (age int, name string) {
	age = 19
	name = "你好"
	return
	// 也可以 return: 19, "Tom" 这样返回
}

// Fun3 多个参数具有相同类型放在一起，可以只写一次类型
func Fun3(a, b, c string, abc, bcd int, p string) (d, e int, g string) {
	d = 15
	e = 16
	g = "你好"
	return
}

// Fun4 不定参数要放在最后面
func Fun4(a string, b int, names ...string) {
	// names就是切片
	for _, name := range names {
		fmt.Printf("不定参数: %s \n", name)
	}
}
