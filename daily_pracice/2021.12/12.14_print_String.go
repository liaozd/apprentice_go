package main

import "fmt"

type T struct {
	a int
	b float64
	c string
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}

// If you want to control the default format for a custom type, all that's required is to define a
// method with the signature String() string on the type. For our simple type T, that might look
// like this.
func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
