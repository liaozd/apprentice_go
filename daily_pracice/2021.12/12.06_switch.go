package main

import "fmt"

func main() {
	ChooseFruit("苹果")
	ChooseFruit("蓝莓")
	ChooseFruit("西噶")
}

func ChooseFruit(fruit string) {
	switch fruit {
	case "苹果":
		fmt.Println("这是苹果\n")
	case "草莓", "蓝莓":
		fmt.Println("这是浆果\n")
	default:
		fmt.Printf("不知道是什么: %s \n", fruit)
	}
}
