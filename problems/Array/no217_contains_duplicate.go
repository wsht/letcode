package arr

import (
	"sort"
)

//这道题是寻找数组中是否有重复数据，
//方法1：暴力模式，双重循环比较 看是否有，
//time O(n^2) space O(1)

//方法二 借助map
//time O(n) space O(n)

//方法三 排序后寻找
//time O(nlog(n)) space O(1)||O(n)

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
