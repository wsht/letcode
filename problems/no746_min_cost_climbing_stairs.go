package main

import (
	"fmt"
	"math"
)

func minCostClimbingStairs(cost []int) int {
	if len(cost) <= 2 {
		return 0
	}
	cost = append(cost, 0)
	return minCostClimbingStairs_help(cost, len(cost)-1)
}

func minCostClimbingStairs_help(cost []int, n int) int {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return cost[n]
	}

	return int(math.Min(float64(minCostClimbingStairs_help(cost, n-1)), float64(minCostClimbingStairs_help(cost, n-2)))) + cost[n]
}

//同问题198 这里也可以将memo 简化为两个变量来进行解答，使空间复杂度为 O(1)
func minCostClimbingStairs2(cost []int) int {
	memo := make([]int, len(cost)+1)
	if len(cost) <= 2 {
		return 0
	}
	memo[0] = cost[0]
	memo[1] = cost[1]
	for i := 1; i < len(cost); i++ {
		tmp_min := int(math.Min(float64(memo[i-1]), float64(memo[i])))
		if i == len(cost)-1 {
			memo[i+1] = tmp_min
		} else {
			memo[i+1] = tmp_min + cost[i+1]
		}
	}
	return memo[len(cost)]
}

func minCostClimbingStairs3(cost []int) int {
	if len(cost) <= 2 {
		return 0
	}
	prev_1 := cost[1]
	prev_2 := cost[0]

	for i := 1; i < len(cost); i++ {
		tmp := prev_1
		tmp_min := int(math.Min(float64(prev_1), float64(prev_2)))
		if i == len(cost)-1 {
			prev_1 = tmp_min
		} else {
			prev_1 = tmp_min + cost[i+1]
		}
		prev_2 = tmp
	}

	return prev_1
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs3([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
