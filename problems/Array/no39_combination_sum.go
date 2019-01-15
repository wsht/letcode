package arr

func CombinationSum(candidates []int, target int) [][]int {
	ret := [][]int{}
	cmbo := []int{}
	backTrack(&ret, cmbo, candidates, 0, target)
	return ret
}

func backTrack(ret *[][]int, cmbo []int, candidates []int, start int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		cmbocp := make([]int, len(cmbo))
		copy(cmbocp, cmbo)
		*ret = append(*ret, cmbocp)
		return
	}

	for i := start; i < len(candidates); i++ {
		val := candidates[i]
		cmbo = append(cmbo, val)
		backTrack(ret, cmbo, candidates, i, target-val)
		cmbo = cmbo[:len(cmbo)-1]
	}
}
