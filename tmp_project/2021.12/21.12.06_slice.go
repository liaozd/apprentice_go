package main

// 陈浩 https://time.geekbang.org/column/article/330178

import (
	"bytes"
	"fmt"
)

func main() {
	fun1()
	fun2()
	fun3()
}

func fun1() {
	fmt.Printf("Func1 子切片")
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100
	fmt.Printf("foo: %v \n", foo)
	// 共享foo的内存地址
	bar := foo[1:4]
	bar[1] = 99
	fmt.Printf("boo: %v \n", bar)
	fmt.Printf("foo: %v \n", foo)
}

func fun2() {
	fmt.Println("\nfun2 append a")
	a := make([]int, 32)
	b := a[1:16]
	fmt.Printf("a address: %p, value: %d \n", a, a)
	fmt.Printf("b address: %p, value: %d \n", b, b)

	// 对 a 做了一个 append()的操作，这个操作会让 a 重新分配内存
	a = append(a, 1)
	fmt.Printf("a address after append: %p value: %d \n", a, a)
}

func fun3() {
	path := []byte("AAA/BBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	// append的"suffix"还没有超出cap，因此不会申请新的地址保存
	dir1 = append(dir1, "suffix"...)
	fmt.Println("after append dir1 with 'suffix")
	fmt.Printf("dir1 => %v, \n", string(dir1))
	fmt.Printf("dir2 => %v \n", string(dir2))
	fmt.Printf("path => %v \n", string(path))

}
