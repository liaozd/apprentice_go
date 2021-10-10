package main

import (
	"fmt"
	"math"
)

func hypot(x float64, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	fmt.Println(hypot(1, 2))
}
