package main

import (
	"fmt"
)

func countBits(num int) []int {
	memo := make([]int, num+1)
	memo[0] = 0
	for i := 1; i <= num; i++ {
		memo[i] = memo[i/2] + i%2
	}

	return memo
}

//位运算中，明显速度比乘除发慢
func countBits2(num int) []int {
	memo := make([]int, num+1)
	memo[0] = 0
	for i := 1; i <= num; i++ {
		memo[i] = memo[i>>1] + i&1
	}

	return memo
}

func main() {
	fmt.Println(countBits(0))
	fmt.Println(countBits(1))
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
	fmt.Println(countBits(6))
	fmt.Println(countBits2(7))
}
