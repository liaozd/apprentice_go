package main

import "fmt"

func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	offset := timeZone["EST"]
	fmt.Println(offset)

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if !attended["person"] {
		fmt.Println("person", "was not at the meeting")
	}

	seconds, ok := timeZone["not exists"]
	fmt.Println(seconds, ok)
}
