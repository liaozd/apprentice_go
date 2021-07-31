package main

import "os"

func main() {
	f, _ := os.Open("xxx.txt")
	f.Close()
}
