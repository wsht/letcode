package arr

func IsToeplitzMatrix(matrix [][]int) bool {
	rowLen := len(matrix)
	if rowLen == 0 {
		return true
	}
	columnLen := len(matrix[0])
	for r := 0; r < rowLen; r++ {
		for c := 0; c < columnLen; c++ {
			if r > 0 && c > 0 && matrix[r-1][c-1] != matrix[r][c] {
				return false
			}
		}
	}
	return true
}
