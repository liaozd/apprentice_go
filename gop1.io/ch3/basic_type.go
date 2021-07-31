package main

import (
	"fmt"
	"math"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	ascii := 'a'
	unicode := '中'
	fmt.Printf("%d %[1]q %[1]c", ascii)
	fmt.Printf("%d %[1]q %[1]c", unicode)

	var f float32 = 16777216
	fmt.Println(f == f+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

}
