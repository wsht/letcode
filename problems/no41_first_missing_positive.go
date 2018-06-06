package main

/**
Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.
*/

//TOOD unfinish
func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] < numsLen && nums[nums[i]-1] != nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}

	for i := 1; i < numsLen; i++ {
		if nums[i] != i {
			return i
		}
	}

	return numsLen + 1
}

func main() {

}
