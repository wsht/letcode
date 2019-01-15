package arr

func ImageSmoother(M [][]int) [][]int {
	r := len(M)
	c := len(M[0])
	ans := [][]int{}

	for i := 0; i < r; i++ {
		ans = append(ans, make([]int, c))
		for j := 0; j < c; j++ {
			count := 0
			for nr := i - 1; nr <= i+1; nr++ {
				for nc := j - 1; nc <= j+1; nc++ {
					if 0 <= nr && nr < r && 0 <= nc && nc < c {
						ans[i][j] += M[nr][nc]
						count++
					}
				}
			}
			ans[i][j] = ans[i][j] / count
		}
	}
	return ans
}
