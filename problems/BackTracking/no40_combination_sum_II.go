package backtrack

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	list := []int{}
	sort.Ints(candidates)
	combinationSum2Helper(&ret, list, candidates, 0, target)
	return ret
}

func combinationSum2Helper(ret *[][]int, list []int, candidates []int, pos, target int) {
	if target == 0 {
		listcp := make([]int, len(list))
		copy(listcp, list)
		*ret = append(*ret, listcp)
		return
	}

	for i := pos; i < len(candidates); i++ {
		// 因为题目中标记了所有的数都是正整数,所以添加此筛选条件
		if target-candidates[i] >= 0 && (i == pos || candidates[i] != candidates[i-1]) {
			list = append(list, candidates[i])
			combinationSum2Helper(ret, list, candidates, i+1, target-candidates[i])
			list = list[:len(list)-1]
		}
	}
}
