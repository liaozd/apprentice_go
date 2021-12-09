package main

import "fmt"

type Member struct {
	id     int
	name   string
	email  string
	gender int
	age    int
}

func main() {
	var m1 Member
	fmt.Printf("%v, %p\n", m1, &m1)

	var m2 = Member{1, "xiaoming", "xiaoming@cn.cn", 1, 1}
	fmt.Printf("%v, %p\n", m2, &m2)

	// 简短变量声明
	m3 := Member{1, "xiaoming", "email", 1, 18}
	fmt.Printf("%v, %p\n", m3, &m3)

	// 其它值为默认"空"值
	var m4 = Member{age: 19, email: "email2"}
	fmt.Printf("%v, %p\n", m4, &m4)

	// 可以访问结构体类型中的字段，或为字段赋值，也可以对字段进行取址(&)操作。
	fmt.Printf("m2.name: %s \n", m2.name)
	fmt.Printf("m2.age value: %d, &m2.age: %p \n", m2.age, &m2.age)

	// * or &
	_age := &m2.age
	*_age = 3000
	fmt.Printf("change m2 age by pointer: m2.age %d, _age: %d, *_age %d\n", m2.age, _age, *_age)
	// *和&是否可以无限套娃
	fmt.Printf("*&m2.age: %d, %p \n", *&m2.age, &*&m2.age)

	// 错误用法，未初始化,nilM1为nil
	var nilM1 *Member
	fmt.Println(nilM1)

	var newM1 = new(Member) // 分配内存来初始化结构休，并返回分配的内存指针，因为已经初始化了，所以可以直接访问字段。
	fmt.Printf("newM1: %v", newM1)

}
