package main

import "fmt"

// Viết hàm có input là 1 string, trả về map có key là ký tự,
// và value là số lần xuất hiện của ký tự đó ở trong string

type CharFrequency map[string]int

func part2() {
	var inputStr string

	fmt.Println("Enter a string:")
	fmt.Scan(&inputStr)

	if inputStr == "" {
		fmt.Println("String cannot be empty")
		return
	}

	fmt.Println("Finished reading input")
	fmt.Println("Input:", inputStr)

	// Create a new CharFrequency map
	charFrequency := CharFrequency{}

	for _, v := range inputStr {
		charFrequency[fmt.Sprintf("%c", v)]++
	}

	fmt.Println("CharFrequency:", charFrequency)
}
