package pkg_name

import (
	"fmt"
	"testing"
)

// 接口就是一个需要实现的方法列表
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct{}

func (g *GoProgrammer) WriteHelloWorld() string {
	fmt.Println("I am GO PROG")
	return "I am go PROG"
}

func TestClient(t *testing.T) {
	var p Programmer // 这样OK
	// p := Programmer  // <- 会报错 Type 'Programmer' is not an expression
	p = new(GoProgrammer)
	p.WriteHelloWorld()
	t.Log(p.WriteHelloWorld())
}
