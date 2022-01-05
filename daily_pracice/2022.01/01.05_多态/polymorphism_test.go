package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgramm(p Programmer) {
	// Programmer是一个interface，必须对应的是指针
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	//goProg := &GoProgrammer{} //  另一种正确方式
	//goProg := GoProgrammer{} // 这样会报错 GoProgrammer does not implement Programmer (WriteHelloWorld method has pointer receiver)

	javaProg := new(JavaProgrammer)
	writeFirstProgramm(goProg)
	writeFirstProgramm(javaProg)
}
