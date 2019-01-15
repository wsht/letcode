package Greedy

import (
	"fmt"
)

func CanJump(nums []int) bool {
	return CanJumpHelper(nums, 0)
}

func CanJumpHelper(nums []int, pos int) bool {
	// fmt.Println(pos)
	if pos > len(nums)-1 {
		return false
	}
	if pos+nums[pos] >= len(nums)-1 {
		fmt.Println(nums[pos], pos)
		return true
	}

	for i := nums[pos]; i >= 1; i-- {
		if CanJumpHelper(nums, pos+i) {
			return true
		}
	}

	return false
}

//我们知道，如果到达不了终点，肯定是因为其中一个点为0，并且，该0点是不可跨越的
//那么只需要知道，该0点前的最最大步数是否大于与零点的距离

//最先我们想到的是 记录该数组中所有的0点位置，然后逐一进行返回计算，
//此方法需要一个 O(n)的space，以及O(N^2)的time

//优化一下，我们美走一步，计算最大跨域，到0点时候，如果最大快跨越大于0点位置，那么则可以跨越

func CanJump2(nums []int) bool {
	n := len(nums)
	max := 0
	//这里只到最后一位即可，所以最后一位不参与计算
	for i := 0; i < n-1; i++ {
		tmp := i + nums[i]
		if tmp > max {
			max = tmp
		}
		if nums[i] == 0 {
			if max <= i {
				return false
			}
		}
	}

	return true
}
