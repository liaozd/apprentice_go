package main

// https://www.jianshu.com/p/431abe0d2ed5?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
// https://www.jianshu.com/p/fc4902159cf5

import (
	"fmt"
)

func main() {
	one(2, callback)
}

func callback(i int) {
	fmt.Printf("I am callback: i%v", i)
}

func one(i int, f func(int)) {
	two(i, fun(f))
}

func two(i int, c Call) {
	c.call(i)
}

type fun func(int)

func (f fun) call(i int) {
	f(i)
}

type Call interface {
	call(int)
}
