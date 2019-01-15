package arr

import (
	"math"
)

// //这个问题其实就是寻找拐点
// func FindUnSortedSubarray(nums []int) int {
// 	start := 0
// 	end := 0
// 	isMatch := false
// 	preIndex := 0
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] < nums[i-1] {
// 			if nums[i] < nums[0] {
// 				start = 0
// 				break
// 			} else if nums[i] == nums[preIndex] {
// 				preIndex++
// 				start = preIndex
// 			} else if !isMatch {
// 				start = i - 1
// 				isMatch = true
// 			}
// 		}
// 	}
// 	preIndex = len(nums) - 1
// 	isMatch = false
// 	for i := len(nums) - 1; i > 0; i-- {
// 		if nums[i] < nums[i-1] {
// 			if nums[i-1] > nums[len(nums)-1] {
// 				end = len(nums) - 1
// 				break
// 			} else if nums[i] == nums[preIndex] {
// 				preIndex--
// 				end = preIndex
// 				// isMatch = true
// 			} else if !isMatch {
// 				end = i
// 				isMatch = true
// 			}
// 		}
// 	}
// 	// fmt.Println(start, end)
// 	if end-start <= 0 {
// 		return 0
// 	}

// 	return end - start + 1
// }

func FindUnSortedSubarray2(nums []int) int {
	n := len(nums)
	min := math.MaxInt32
	max := math.MinInt32
	flag := false
	//从左到右 找到第一拐点以后的最小值
	for i := 0; i < n; i++ {
		if nums[i] < nums[i-1] {
			flag = true
		}
		if flag {
			min = int(math.Min(float64(min), float64(nums[i])))
		}
	}

	// 从右到左找到第一个拐点后的最大值
	flag = false
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			flag = true
		}
		if flag {
			max = int(math.Max(float64(max), float64(nums[i])))
		}
	}
	//从左到右找到比最小值更小的点
	left, right := 0, 0
	for left = 0; left < n; left++ {
		if min < nums[left] {
			break
		}
	}
	//从右到左找到比最大值更大的点
	for right = n - 1; right >= 0; right-- {
		if max > nums[right] {
			break
		}
	}

	if right-left >= 0 {
		return right - left + 1
	}

	return 0
}
