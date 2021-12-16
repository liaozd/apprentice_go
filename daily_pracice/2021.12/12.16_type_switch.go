package main

import "fmt"

type Stringer interface {
	String() string
}

func main() {
	fmt.Println(switchType())
}

func switchType() string {
	var value interface{}
	switch str := value.(type) {
	case string:
		return str // 获取具体的值
	case Stringer:
		return str.String() // 将该接口转换为另一个接口
	}
	return ""
}
