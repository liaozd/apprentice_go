package main

import (
	"fmt"
	"reflect"
)

type Data struct{}

func main() {
	v1 := Data{}
	v2 := Data{}
	fmt.Println("v1 == v2", reflect.DeepEqual(v1, v2))

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"one": "a", "two": "b"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
}
