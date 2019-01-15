package arr

import (
	"fmt"
)

func ProductExceptSelf(nums []int) []int {
	//find 0
	ret := make([]int, len(nums))
	zerocount := 0
	zeroIndex := -1
	pro := 1
	next := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zerocount++
			zeroIndex = i
		} else {
			pro *= nums[i]
			if i > 1 {
				next *= nums[i]
			}
		}
		if zerocount >= 2 {
			return ret
		}
	}

	fmt.Println(pro, next)
	if zerocount == 1 {
		ret[zeroIndex] = pro
	} else {
		for i := 0; i < len(nums); i++ {
			ret[i] = pro / nums[i]
		}
	}

	return ret
}

//这你吗，是怎么想出来的，不用除法就可以做成这样
func ProductExceptSelf2(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
		fmt.Println(res)
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
		// fmt.Println(res, right)
	}
	return res
}
