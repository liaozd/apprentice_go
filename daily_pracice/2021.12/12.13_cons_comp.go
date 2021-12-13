package main

import "fmt"

func main() {
	f := NewFile(1, "aa")
	fmt.Println(f)
}

//
func NewFile1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.nepipe = 0
	return f
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0}
	//返回一个局部变量的地址完全没有问题，这点与 C 不同。该局部变量对应的数据 在
	//函数返回后依然有效。实际上，每当获取一个复合字面的地址时，都将为一个新的实例分配
	//内存

	return &f
}
