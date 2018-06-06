package main

/**
右转90度 可以先水平对称，然后斜线对称
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	//水平对称
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}
	//按照左中心线对称
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
