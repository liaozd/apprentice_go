package main

import "fmt"

func replace(mySlice []string) {
	mySlice[0] = "nima"
}
func main() {
	// 切片slice
	//var courses []string // 定义一个切片
	//courses := []string{"d", "s", "t"}
	//fmt.Printf("%T", courses)

	// make初始化切片，第二个参数是5切片的长度
	m_courses := make([]string, 5)
	fmt.Printf("%T", m_courses)

	// 第三种方法：通过数组变成一个切片
	var courses = [5]string{"d", "s", "t"}
	subCourse := courses[1:4]
	fmt.Printf("%T", subCourse)

	// 第四种：new
	new_subCourse := *new([]int)
	fmt.Printf("%T", new_subCourse)

	// 数组的传递是值传递，切片是引用传递、不是值传递
	replace(subCourse)
	fmt.Println(subCourse)

}
