package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	the_max, the_min := prices[0], prices[0]
	max_profit := 0
	for i := 1; i < len(prices); i++ {
		if the_min > prices[i] {
			the_min = prices[i]
			the_max = prices[i]
		} else if the_max < prices[i] {
			the_max = prices[i]
			if the_max-the_min > max_profit {
				max_profit = the_max - the_min
			}
		}
	}
	return max_profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
