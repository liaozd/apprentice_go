package main

import "fmt"

func main() {

	//var courses [10] string
	//var course = [5]string{"a", "b", "c"}
	course := [5]string{"a", "b", "c"}
	//var course [3] string = [3]string{"a", "b", "c"}
	course[0] = "4"
	fmt.Println(course)

	var a = [4]float32{}
	fmt.Println(a)

	var c = [5]int{'A', 'B'}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	e := [5]int{4: 100}
	fmt.Println(e)

	f := [...]int{0: 1, 4: 1, 9: 100}
	fmt.Println(f)

	// 数组长度
	fmt.Println(len(f))

	// 便利
	for i, value := range course {
		fmt.Println(i, value)
	}

	sum := 0
	for _, value := range f {
		sum += value
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(course); i++ {
		sum += f[i]
		fmt.Println(f[i])
	}
	fmt.Println(sum)

	// 数组是值类型，数组长度不一样，类型不同
	courseA := [3]string{"django", "scrapy", "tornado"}
	courseB := [...]string{"django1", "s", "t", "s"}
	fmt.Printf("%T\n", courseA)
	fmt.Printf("%T\n", courseB)
}
