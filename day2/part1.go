package main

import (
	"fmt"
	"time"
)

// Viết hàm tạo ra 1 struct về 1 người (gồm tên, nghề nghiệp, năm sinh) ,
//  và struct có method tính tuổi,
// method kiểm tra người ấy có hợp với nghề của mình không
// ( năm sinh chia hết cho số chữ trong tên)

type Person struct {
	Name      string
	Job       string
	BirthYear int
}

func (p *Person) currentAge() int {
	currentYear := time.Now().Year()
	return currentYear - p.BirthYear
}

func (p *Person) isMatchingJob() bool {
	numberOfCharactersInName := len(p.Name)

	return p.BirthYear%numberOfCharactersInName == 0
}

func part1() {
	var name, job string
	var birthYear int

	fmt.Println("Enter Name:")
	fmt.Scan(&name)

	fmt.Println("Enter Job:")
	fmt.Scan(&job)

	fmt.Println("Enter Birth Year:")
	fmt.Scan(&birthYear)

	fmt.Println("Name:", name)
	fmt.Println("Job:", job)
	fmt.Println("Birth Year:", birthYear)

	// Create a new person
	p := Person{
		Name:      name,
		Job:       job,
		BirthYear: birthYear,
	}

	if p.Name == "" || p.Job == "" || p.BirthYear == 0 {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println("FINISH READING INFO")
	fmt.Println("Info of person:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Job:", p.Job)
	fmt.Println("Birth Year:", p.BirthYear)
	fmt.Println("Current Age:", p.currentAge())
	fmt.Println("Is Matching Job:", p.isMatchingJob())
}
