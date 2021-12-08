package main

import "fmt"

func main() {
	name := "Tom"
	age := 17

	// 这个 API 是返回字符串的，所以大多数时候我们都是用这个
	str := fmt.Sprintf("Hello, I am %s, I am %d years old \n", name, age)
	println(str)

	// 直接输出，DEBUG用
	fmt.Printf("Hello, I am %s, I am %d years old \n", name, age)

	replaceHolder()
}

func replaceHolder() {
	u := &user{
		Name: "Tom",
		Age:  17,
	}
	fmt.Printf("v => %v \n", u)
	fmt.Printf("u address => %p \n", u)
	fmt.Printf("+v => %+v \n", u)
	fmt.Printf("#v => %#v \n", u)
	fmt.Printf("T => %T \n", u)
}

type user struct {
	Name string
	Age  int
}
