package main

import (
	"fmt"
)

func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	fmt.Printf("Hello %d\n", 23)

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	fmt.Printf("%v\n", timeZone)

	type T struct {
		a int
		b float64
		c string
	}

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("T struct: %v\n", t)
	fmt.Printf("T struct %%+v: %+v\n", t)
	fmt.Printf("T struct %%#v: %#v\n", t)
	fmt.Printf("timeZone: %#v\n", timeZone)
}
