package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x //取地址

	fmt.Println(p)
	fmt.Println(*p) //从地址中取值

	*p = 2
	fmt.Println(*p)
	fmt.Println(p)

	v := 1
	incr(&v)              //v现在是2
	fmt.Println(incr(&v)) //v现在是3

}

func incr(p *int) int {
	*p++ //p是局部变量，但是*p代表函数外的变量
	return *p
}
