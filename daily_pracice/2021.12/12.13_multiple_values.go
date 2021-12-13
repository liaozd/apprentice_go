package main

import (
	"fmt"
	"unicode"
)

func main() {
	var x int
	b := "abcd1234"
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !unicode.IsDigit(rune(b[i])); i++ {
		fmt.Printf("%v is not digit\n", b[i])
	}
	x := 0
	for ; i < len(b) && unicode.IsDigit(rune(b[i])); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}
