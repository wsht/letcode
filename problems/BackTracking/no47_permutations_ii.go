package backtrack

import (
	"sort"
)

func PermuteUnique(nums []int) [][]int {
	ret := [][]int{}
	used := make([]bool, len(nums))
	list := []int{}
	sort.Ints(nums)
	PermuteUniqueHelper(&ret, used, nums, list)
	return ret
}

func PermuteUniqueHelper(ret *[][]int, used []bool, nums, list []int) {
	if len(list) == len(nums) {
		listcp := make([]int, len(list))
		copy(listcp, list)
		*ret = append(*ret, listcp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 1,1,2 的时候， 当前·位置为1， 只有0被使用的时候，我们才可以继续
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		list = append(list, nums[i])
		PermuteUniqueHelper(ret, used, nums, list)
		used[i] = false
		list = list[:len(list)-1]

	}
}
