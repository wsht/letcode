package main

import "fmt"

//我们可以知道 if n>0 climb(n)  = climb(n-1) + climb(n-2) + 1
func climbStair(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	return climbStair(n-1) + climbStair(n-2)
}

//存储已经计算的结果
func climbStair2(n int) int {
	memo := make([]int, n+1)
	return climbStair2_help(n, memo)
}

func climbStair2_help(n int, memo []int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return n
	}

	if memo[n] > 0 {
		return memo[n]
	}

	result := climbStair2_help(n-1, memo) + climbStair2_help(n-2, memo)
	memo[n] = result
	return result
}

func climbStair3(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	memo := make([]int, n+1)
	memo[1] = 1
	memo[2] = 2
	for i := 3; i < n+1; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	fmt.Println(memo)
	return memo[n]
}

func climbStair4(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	floor_n1 := 1
	floor_n2 := 2
	for i := 3; i < n+1; i++ {
		tmp := floor_n2
		floor_n2 = floor_n2 + floor_n1
		floor_n1 = tmp
	}
	return floor_n2
}

func main() {
	fmt.Println(climbStair4(1))
	fmt.Println(climbStair4(2))
	fmt.Println(climbStair4(4))
}
