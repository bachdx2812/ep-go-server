package main

import "fmt"

func part2() {
	var inputStr string

	fmt.Println("Enter a string:")
	fmt.Scan(&inputStr)

	if inputStr == "" {
		fmt.Println("String cannot be empty")
		return
	}

	if len(inputStr)%2 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
