package main

import "fmt"

func main() {
	score := 85

	grade := "A"
	//if score > 90 {
	//	grade = "A"
	//} else if score > 80 && score < 89 {
	//	grade = "B"
	//} else {
	//	grade = "E"
	//}

	switch {
	case score > 90:
		grade = "A"
		fallthrough
	case score > 80 && score <= 90:
		grade = "B"
	default:
		grade = "default"
	}
	fmt.Println(grade)

	var a = "daddy"
	switch a {
	case "mum", "daddy":
		fmt.Println(a)
	}
}
