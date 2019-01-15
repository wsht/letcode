package arr

import (
	"math"
)

func FindShortestSubArrayWithHashMap(nums []int) int {
	countMap := map[int]int{}
	left := map[int]int{}
	right := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := left[nums[i]]
		if !ok {
			left[nums[i]] = i
		}
		right[nums[i]] = i
		countMap[nums[i]] += 1
	}
	ans := len(nums)
	degree := 0
	for one := range countMap {
		if countMap[one] > degree {
			degree = countMap[one]
			ans = right[one] - left[one]
		} else if countMap[one] == degree {
			ans = int(math.Min(float64(right[one]-left[one]), float64(ans)))
		}
	}
	return ans + 1
}
