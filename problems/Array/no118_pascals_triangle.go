package arr

func Generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		res = append(res, make([]int, i+1))
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
