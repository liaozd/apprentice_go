package main

import (
	"fmt"
	"os"
)

var (
	home   = os.Getenv("HOME")
	user1  = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func main() {
	fmt.Println(home)
	fmt.Println(user1)
	fmt.Println(gopath)
}
