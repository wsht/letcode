package backtrack

import (
	"sort"
	"strconv"
)

func SubSetWithDup(nums []int) [][]int {
	ret := [][]int{}
	list := []int{}
	vmap := map[string]int{}
	sort.Ints(nums)
	SubSetWithDupHelper(&ret, nums, list, "", vmap, 0)
	return ret
}

func SubSetWithDupHelper(ret *[][]int, nums []int, list []int, pre string, vmap map[string]int, pos int) {

	listcp := make([]int, len(list))
	copy(listcp, list)
	*ret = append(*ret, listcp)

	for i := pos; i < len(nums); i++ {
		key := pre + strconv.Itoa(nums[i])
		_, ok := vmap[key]
		if !ok {
			vmap[key] = 1
			list = append(list, nums[i])
			SubSetWithDupHelper(ret, nums, list, key, vmap, i+1)
			list = list[:len(list)-1]
			// vmap[key] = 0
		}
	}
}

func SubSetWithDup2(nums []int) [][]int {
	ret := [][]int{}
	list := []int{}
	sort.Ints(nums)
	SubSetWithDupHelper2(&ret, nums, list, 0)
	return ret
}

func SubSetWithDupHelper2(ret *[][]int, nums []int, list []int, pos int) {
	listcp := make([]int, len(list))
	copy(listcp, list)
	*ret = append(*ret, listcp)

	for i := pos; i < len(nums); i++ {
		if i == pos || nums[i] != nums[i-1] {
			list = append(list, nums[i])
			SubSetWithDupHelper2(ret, nums, list, i+1)
			list = list[:len(list)-1]
		}
	}
}
