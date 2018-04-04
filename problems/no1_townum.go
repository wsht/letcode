package main

import (
	"fmt"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

/**
time O(n^2)
space O(1)
但是这里会有重复的轮询
*/
func twoSum(nums []int, target int) []int {
	var result = make([]int, 2)

	for i, iValue := range nums {
		for j, jValue := range nums {
			if i != j && target == iValue+jValue {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

/**
方案1的优化 已经验证过的将不再验证
*/
func twoSumSolution2(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i + i; j < numsLen; j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

/**
Two pass hash table
time complexity O(n)
space complexity O(n)
*/
func twoSumSolution3(nums []int, target int) []int {
	maps := make(map[int]int)
	fmt.Println(maps[1])
	for i, value := range nums {
		maps[value] = i
	}

	for i, value := range nums {
		complement := target - value
		index, ok := maps[complement]
		if ok && index != i {
			return []int{i, index}
		}
	}

	return []int{}
}

/**
One pass hash table
只执行一次循环

time complexity O(n)
space complexity O(n)
*/
func towSumSolution4(nums []int, target int) []int {
	maps := make(map[int]int)

	for i, value := range nums {
		complement := target - value
		index, ok := maps[complement]
		if ok && index != i {
			return []int{i, index}
		}
		maps[value] = i
	}

	return []int{}
}

// func main() {
// 	input := []int{3, 2, 4}

// 	fmt.Println(twoSum(input, 6))
// }
