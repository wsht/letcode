package main

import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon[0])
	n := len(dungeon)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i] = append(dp[i], math.MaxInt64)
		}
	}
	dp[n][m-1] = 1
	dp[n-1][m] = 1

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			need := int(math.Min(float64(dp[i+1][j]), float64(dp[i][j+1]))) - dungeon[i][j]
			if need <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = need
			}
		}
	}
	return dp[0][0]
}

// func main() {
// 	m, n := 3, 3
// 	fmt.Println(m)
// 	dp := make([][]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		for j := 0; j < m+1; j++ {
// 			dp[i] = append(dp[i], math.MaxInt64)
// 		}
// 	}
// 	fmt.Println(dp)
// }
