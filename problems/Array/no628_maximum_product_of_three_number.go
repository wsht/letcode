package arr

import (
	"fmt"
	"math"
)

func MaximumProduct(nums []int) int {
	max := [3]int{math.MinInt32, math.MinInt32, math.MinInt32}
	min := [2]int{math.MaxInt32, math.MaxInt32}
	for i := 0; i < len(nums); i++ {
		max, min = RebuildResult(max, min, nums[i])
	}
	fmt.Println(max, min)
	return int(math.Max(float64(min[0]*min[1]*max[2]), float64(max[0]*max[1]*max[2])))
}

func RebuildResult(max [3]int, min [2]int, value int) ([3]int, [2]int) {
	tmpValue := value
	for i := 2; i >= 0; i-- {
		tmp := max[i]
		if tmp < tmpValue {
			max[i] = tmpValue
			tmpValue = tmp
		}
	}
	tmpValue = value
	for i := 0; i <= 1; i++ {
		tmp := min[i]
		if tmp > tmpValue {
			min[i] = tmpValue
			tmpValue = tmp
		}
	}
	return max, min
}
