package main

import "fmt"

func main() {
	maked := make([]int, 10, 100)
	newed := new([]int)

	fmt.Println(maked)
	fmt.Println(newed)

	var p *[]int = new([]int)    // 分配切片结构，new返回一个指针
	var v []int = make([]int, 5) // 切片，v引用了一个100个0 int元素的新数组
	// make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用
	//new 分配内存。
	fmt.Printf("new: %v", p)
	fmt.Printf("make: %v", v)
}
