package main

import (
	"fmt"
	"unsafe"
)

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) changeBase(f float64) {
	// 方法changeBase接受的是一个值引用。这意味着这个方法操作的是结构体Triangle的副本，
	// 而原始结构体不受影响。在方法changeBase中，修改的是原始Triangle结构体的副本的t.base。
	fmt.Println("t pointer in main: ", unsafe.Pointer(&t))
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 1}
	fmt.Println("t pointer in main: ", unsafe.Pointer(&t))
	t.changeBase(4)
	fmt.Println(t.base)
}
