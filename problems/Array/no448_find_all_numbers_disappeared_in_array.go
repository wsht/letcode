package arr

import (
	"math"
)

func FindDisappearedNumbers(nums []int) []int {
	result := []int{}
	//因为我们知道这个是等差数列，那么用值作为标记，将不重复的赋值为负数
	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	//寻找非0的 下标的 i , i+1就是结果值
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
