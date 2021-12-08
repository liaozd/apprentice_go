package main

import (
	"flag"
	"fmt"
)

var cliName = flag.String("name", "nick", "Input Your Name")
var cliAge = flag.Int("age", 28, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
}

func main() {
	Init()
	flag.Parse()

	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
}
