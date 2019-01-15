package main

import (
	"fmt"
	"math"
)

func superEggDrop(K, N int) int {
	dp := [][]int{}
	for i := 0; i <= K; i++ {
		dp = append(dp, make([]int, N+1))
	}
	result := helper(K, N, dp)
	fmt.Println(dp)
	return result
	// if K >= 1 {
	// 	for i := 0; i <= N; i++ {
	// 		dp[1][i] = i
	// 	}
	// }

	// for i := 2; i <= K; i++ {
	// 	for j := 1; j <= N; j++ {
	// 		mid := j / 2
	// 		if mid == 0 {
	// 			mid = 1
	// 		}
	// 		tmp := 1 + int(math.Max(float64(dp[i-1][j-1]), float64(dp[i][N-j])))
	// 		fmt.Println(tmp)
	// 		min := j
	// 		dp[i][j] = int(math.Min(float64(min), float64(tmp)))
	// 	}
	// }
	// fmt.Println(dp)
	// return dp[K][N]

}

func helper(K, N int, dp [][]int) int {
	// fmt.Println(K, N)
	if N <= 1 {
		return N
	}
	if K == 1 {
		return N
	}

	if dp[K][N] > 0 {
		return dp[K][N]
	}

	min := N
	for i := 1; i <= N; i++ {
		left := helper(K-1, i-1, dp)
		right := helper(K, N-i, dp)
		min = int(math.Min(float64(min), float64(1)+math.Max(float64(left), float64(right))))
	}
	dp[K][N] = min
	return min
}

func superEggDrop2(K, N int) int {
	dp := make([]int, K+1)
	m := 0
	for m = 0; dp[k] < N; m++ {
		for k := K; k > 0; k-- {
			dp[k] += dp[k-1] + 1
		}
	}
	return m
}

func main() {
	superEggDrop(5, 10000)
}
