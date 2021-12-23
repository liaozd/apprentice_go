package main

import "fmt"

type Person_ struct {
	Name string
	Age  int
}

func ValueReceiver(p Person_) {
	p.Name = "John"
	fmt.Println("Inside ValueReceiver : ", p.Name)
}

func PointerReceiver(p *Person_) {
	p.Age = 24
	fmt.Println("Inside PointerReceiver model: ", p.Age)
}

func main() {
	p := Person_{"Tom", 28}
	p1 := &Person_{"Patric", 68}

	ValueReceiver(p)
	fmt.Println("Inside Main after value receiver : ", p.Name)

	fmt.Println("Inside Main before value receiver : ", p1.Age)
	PointerReceiver(p1)
	fmt.Println("Inside Main after value receiver : ", p1.Age)
}
