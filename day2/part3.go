package main

import "fmt"

// Two sum: https://leetcode.com/problems/two-sum/
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

func inputSlice(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return inputSlice(x, err)
}

func part3() {
	fmt.Println("Enter input:")
	x := inputSlice([]int{}, nil)
	fmt.Println("Finished Input:", x)

	var target int
	fmt.Println("Enter target:")
	fmt.Scan(&target)
	fmt.Println("Target:", target)

	type MapAppearance map[int]int
	mapAppearance := MapAppearance{}

	for index, value := range x {
		// find in mapAppearance
		if idx, ok := mapAppearance[target-value]; ok {
			fmt.Printf("Found at [%d, %d]\n", index, idx)
			return
		}

		mapAppearance[value] = index
	}
}
