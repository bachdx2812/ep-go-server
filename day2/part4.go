package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Human struct {
	Name       string
	Occupation string
	BirthYear  int
}

func part4() {
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var humans []Human

	// Create a new buffered reader
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into parts by the delimiter '|'
		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		// Parse the birth year as an integer
		birthYear, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error parsing birth year:", err)
			continue
		}

		human := Human{
			Name:       strings.ToUpper(parts[0]),
			Occupation: strings.ToLower(parts[1]),
			BirthYear:  birthYear,
		}
		humans = append(humans, human)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, h := range humans {
		fmt.Printf("%+v\n", h)
	}
}
