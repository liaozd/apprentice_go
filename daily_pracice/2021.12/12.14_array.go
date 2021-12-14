package main

import "fmt"

func main() {
	array := [...]float64{7.0, 8.5, 9}
	x := Sum(&array)
	fmt.Println(x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}
