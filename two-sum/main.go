package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// map
	numMap := make(map[int]int)
	// slice, array is fixed size?
	var result []int

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		// Check if the complement exists in the map
		// If it does, j will be the index of the complement, and ok will be true
		if j, ok := numMap[complement]; ok {
			// If the complement is found, append the indices to the result
			result = append(result, j, i)
			// Return the result as we found the solution
			return result
		}

		// If the complement is not found, add the current number and its index to the map
		numMap[nums[i]] = i
	}

	// If no solution is found, return an empty slice
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
