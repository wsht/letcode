package backtrack

func Combine(n, k int) [][]int {
	ret := [][]int{}
	list := []int{}
	CombineHelper(&ret, list, 1, n, k)
	return ret
}

func CombineHelper(ret *[][]int, list []int, start, n, k int) {
	if len(list) == k {
		listcp := make([]int, k)
		copy(listcp, list)
		*ret = append(*ret, listcp)
		return
	}

	for i := start; i <= n; i++ {
		list = append(list, i)
		CombineHelper(ret, list, i+1, n, k)
		list = list[:len(list)-1]
	}
}
