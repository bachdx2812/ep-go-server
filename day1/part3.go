package main

import (
	"fmt"
	"sort"
)

func inputForPart3(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return inputForPart3(x, err)
}

func part3() {
	fmt.Println("Enter input:")
	x := inputForPart3([]int{}, nil)
	fmt.Println("Finished Input:", x)

	// SUM of input
	var sum int

	for _, v := range x {
		sum += v
	}
	fmt.Println("Sum of input:", sum)
	// END SUM

	// AVG of input
	avg := float64(sum) / float64(len(x))
	fmt.Println("Average of input:", avg)
	// END AVG

	// MAX of input
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}

	fmt.Println("Max of input:", max)
	// END MAX

	// MIN of input
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	fmt.Println("Min of input:", min)
	// END MIN

	// Sort input
	sort.Ints(x)
	fmt.Println("Sorted input:", x)
	// END SORT
}
