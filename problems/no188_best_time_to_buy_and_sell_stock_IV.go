package main

import (
	"fmt"
	"math"
)

func maxProfit(k int, prices []int) int {
	day := len(prices)
	dp := [][]int{}
	for i := 0; i < day; i++ {
		dp = append(dp, make([]int, day))
		for j := 0; j < day; j++ {
			dp[i][j] = -1
		}
	}
	//这里直接构造？
	for i := 0; i < day; i++ {
		max := 0
		for j := i + 1; j < day; j++ {
			income := prices[j] - prices[i]
			if income >= 0 {
				dp[i][j] = income
			} else {
				dp[i][j] = 0
			}
			if income > max {
				max = income
			}
			dp[j][i] = max
		}
	}

	fmt.Println(dp)
	// return 1
	return helper(k, 0, day, dp)
}

func helper(k int, day int, max_day int, dp [][]int) int {
	max := 0
	// fmt.Println(k)
	if k <= 0 {
		return max
	}
	for i := day; i < max_day; i++ {
		for j := i + 1; j < max_day; j++ {
			tmp := helper(k-1, j, max_day, dp)
			tmp_max := dp[j][i] + tmp
			if tmp_max > max {
				max = tmp_max
			}
		}
	}
	// dp[max_day-1][day] = max
	return max
}

func maxProfit2(k int, prices []int) int {
	prices_len := len(prices)
	if prices_len < 2 {
		return 0
	}
	if k > prices_len/2 {
		ans := 0
		for i := 1; i < prices_len; i++ {
			ans += int(math.Max(float64(prices[i]-prices[i-1]), 0))
		}
		return ans
	}
	hold := make([]int, k+1)
	rele := make([]int, k+1)
	for i := 0; i <= k; i++ {
		hold[i] = math.MinInt64
		rele[i] = 0
	}

	cur := 0
	for i := 0; i < prices_len; i++ {
		cur = prices[i]
		for j := k; j > 0; j-- {
			rele[j] = int(math.Max(float64(rele[j]), float64(hold[j]+cur)))
			hold[j] = int(math.Max(float64(hold[j]), float64(rele[j-1]-cur)))
		}
	}

	return rele[k]
}

func main() {
	// fmt.Println(maxProfit(7, []int{48, 12, 60, 93, 97, 42, 25, 64, 17, 56, 85, 93, 9, 48, 52, 42, 58, 85, 81, 84, 69, 36, 1, 54, 23, 15, 72, 15, 11, 94}))
	fmt.Println(maxProfit(1, []int{6, 1, 6, 4, 3, 0, 2}))
	// 1

}
