package main

import (
	"fmt"
	"math"
)

//我们可以推测出 rob(i) = max((rob(i-2) + cur_value), rob(i-1))
func rob(nums []int) int {
	return rob_helper(len(nums)-1, nums)
}
func rob_helper(i int, nums []int) int {
	if i <= 0 {
		return 0
	}
	return int(math.Max(float64(rob_helper(i-2, nums)+nums[i]), float64(rob_helper(i-1, nums))))
}

//现在我们引入 memo 用来存储已经计算过的值 //top-down
func rob2(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = -1
	}
	return rob2_helper(nums, dp, len(nums)-1)
}
func rob2_helper(nums []int, dp []int, i int) int {
	if i <= 0 {
		return 0
	}

	if dp[i] >= 0 {
		return dp[i]
	}

	result := int(math.Max(float64(rob2_helper(nums, dp, i-2)+nums[i]), float64(rob2_helper(nums, dp, i-1))))
	dp[i] = result
	return result
}

//bottom-up
func rob3(nums []int) int {
	dp := make([]int, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		val := nums[i]
		dp[i+1] = int(math.Max(float64(dp[i]), float64(dp[i-1]+val)))
	}

	return dp[len(nums)]
}

//我们可以注意到， 上个方法中，我们只用到了dp[i] 以及dp[i-1]，所以，我们可以只用两个变量来存储他
func rob4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev_1, prev_2 := 0, 0
	for _, num := range nums {
		tmp := prev_1
		prev_1 = int(math.Max(float64(prev_2+num), float64(prev_1)))
		prev_2 = tmp
	}

	return prev_1
}

func main() {
	fmt.Println(rob([]int{1, 1, 1, 2}))
}
