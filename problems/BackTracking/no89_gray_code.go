package backtrack

func GrayCode(n int) []int {
	ret := []int{}
	for i := 0; i < 1<<uint(n); i++ {
		ret = append(ret, i^i>>1)
	}
	return ret
}
