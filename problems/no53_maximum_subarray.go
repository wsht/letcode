package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	max := math.MinInt64
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = int(math.Max(float64(sum), 0)) + nums[i]
		max = int(math.Max(float64(max), float64(sum)))
	}
	return max
}
