package arr

import (
	"math"
)

func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	max := 0
	if n <= k {
		max = nums[n-1]
		k = n
	} else {
		max = nums[k-1]
	}
	for i := k; i < n; i++ {
		max = int(math.Max(float64(max), float64(nums[i]-nums[i-k])))
	}

	return float64(max) / float64(k)
}

//sliding window

func FindMaxAverageSlidingWidnow(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	res := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		res = int(math.Max(float64(sum), float64(res)))
	}
	return float64(res) / float64(k)
}
