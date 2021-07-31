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

	// slice是动态数组，可以动态添加值
	fmt.Println(subCourse[1])
	subCourse2 := subCourse[1:3]
	fmt.Printf("%T, %v", subCourse2, subCourse2)

	// append可以向切片追加元素
	subCourse2 = append(subCourse2, "imooc")
	fmt.Println(subCourse2)

	// 初始化0长度的切片，只将0长度值copy进去
	subCourse3 := []string{}
	fmt.Println(len(subCourse3))
	// 拷贝的时候目标对象长度需要设置好保持一致
	copy(subCourse3, subCourse2)
	fmt.Println(subCourse3)

}
