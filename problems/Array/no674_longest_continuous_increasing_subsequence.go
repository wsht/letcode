package arr

import (
	"math"
)

func FindLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			max = int(math.Max(float64(max), float64(count)))
			count = 1
		}
	}
	max = int(math.Max(float64(max), float64(count)))
	return max
}
