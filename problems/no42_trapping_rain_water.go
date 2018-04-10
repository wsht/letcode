package main

import (
	"math"
)

func trap2(height []int) int {
	arrLen := len(height)
	result := 0
	for i := 1; i < arrLen-1; i++ {
		curValue := height[i]
		for trap2Compare(height, curValue, i) {
			curValue += 1
			result += 1
		}
	}

	return result
}

func trap2Compare(height []int, value, index int) bool {
	left := index - 1
	right := index + 1
	leftResult := false
	rightResult := false
	for left >= 0 {
		if height[left] > value {
			leftResult = true
		}
		left--
	}

	for right < len(height) {
		if height[right] > value {
			rightResult = true
		}
		right++
	}

	return leftResult && rightResult
}

/**
改进版本，找到最小的比当前高度高的bar 进行累加并且排除自身高度
*/

func trap3(height []int) int {
	ans := 0
	size := len(height)
	for i := 1; i < size-1; i++ {
		maxLeft, maxRight := 0, 0
		for j := i; j >= 0; j-- {
			maxLeft = int(math.Max(float64(maxLeft), float64(height[j])))
		}
		for j := i; j < size; j++ {
			maxRight = int(math.Max(float64(maxRight), float64(height[j])))
		}

		ans += int(math.Min(float64(maxLeft), float64(maxRight))) - height[i]
		// fmt.Println(ans)
	}

	return ans
}

/**
Approach #4 using two pointers
*/
func trap4(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}

	return ans
}

// func main() {
// 	fmt.Println(trap3([]int{1, 2, 3, 4}))
// }
