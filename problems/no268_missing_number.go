package main

import (
	"fmt"
)

// func missingNUmber(nums []int) int {
// 	numsLen := len(nums)
// 	min := 0
// 	for i := 0; i < numsLen; i++ {
// 		if numsLen-nums[i] < min {
// 			// max = numsLen - nums[i]
// 		}
// 	}

// 	return max
// }

func missingNumber(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}

	return missing
}

func main() {
	fmt.Println(4 ^ 3 ^ 2)
}
