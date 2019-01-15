package main

import (
	"fmt"
)

func numberOfArithmeticSlices(A []int) int {

	dp := make([]int, 4)
	dp[3] = 1
	if len(A) <= 3 {
		return helper(len(A), dp)
	}
	result_num := 0
	max_arithmetic_num := 0
	for i := 0; i < len(A)-2; i++ {

		if A[i]+A[i+2] == A[i+1]*2 {

			if max_arithmetic_num == 0 {
				max_arithmetic_num = 3
			} else {
				max_arithmetic_num += 1
			}
		} else {

			result_num += helper(max_arithmetic_num, dp)
			max_arithmetic_num = 0
		}
	}
	result_num += helper(max_arithmetic_num, dp)

	return result_num
}

func helper(n int, dp []int) int {
	if n < len(dp) {
		return dp[n]
	} else {
		for i := len(dp) - 1; i < n; i++ {
			dp = append(dp, i+1-2+dp[i])
		}
	}
	return dp[n]
}

func main() {
	// dp := make([]int, 4)
	// dp[3] = 1
	// fmt.Println(helper(4, dp))
	// fmt.Println(helper(7, dp))
	// for i := 0; i < 100; i++ {
	// 	dp = append(dp, 1)
	// }
	// fmt.Println(dp)
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 7}))
}
