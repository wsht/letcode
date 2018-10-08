package main

func uniquePaths(m, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[1][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j]
		}
	}

	return dp[n][m]
}

// func main() {
// 	fmt.Println(uniquePaths(7, 3))
// }
