package arr

import (
	"math"
)

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	consecutive := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			consecutive++
		} else {
			if consecutive != 0 {
				max = int(math.Max(float64(max), float64(consecutive)))
				consecutive = 0
			}
		}
	}
	max = int(math.Max(float64(max), float64(consecutive)))

	return max
}
