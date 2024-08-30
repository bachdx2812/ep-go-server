package main

import "fmt"

func part1() {
	var length, width float32

	fmt.Println("Enter Length:")
	fmt.Scan(&length)

	fmt.Println("Enter Width:")
	fmt.Scan(&width)

	// check if length is valid
	if length <= 0 || width <= 0 {
		fmt.Println("Length and Width must be greater than 0")
		return
	}

	fmt.Printf("The perimeter of the rectangle is %.2f\n", 2*(length+width))
	fmt.Printf("The area of the rectangle is %.2f\n", length*width)
}
