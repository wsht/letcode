package main

import (
	"fmt"
	"math"
	"strconv"
)

func atMostNGivenDigitSet(D []string, N int) int {
	fmt.Println(D)
	dp := make([]int, 10)
	min, err := strconv.Atoi(D[0])
	max, err := strconv.Atoi(D[len(D)-1])
	if err != nil {

	}
	tmp := N
	min_nums := 0
	max_nums := 0
	n_len := 0
	for i := 0; tmp > 0; i++ {
		min_nums += min * int(math.Pow10(i))

		max_nums += max * int(math.Pow10(i))
		tmp = tmp / 10
		n_len = i
	}
	fmt.Println("nlen", n_len)

	fmt.Println(min_nums, max_nums)
	for i := 0; i < len(D); i++ {
		num, err := strconv.Atoi(D[i])
		if err != nil {

		}
		for j := num; j <= 9; j++ {
			dp[j] = i + 1
		}
	}

	if N < min_nums {
		N = max_nums / 10
	} else {
		index := N / int(math.Pow10(n_len))
		if dp[index-1] == dp[index] {
			num, err := strconv.Atoi(D[dp[index]-1])
			if err != nil {

			}
			N = max_nums%int(math.Pow10(n_len)) + num*int(math.Pow10(n_len))
		}
	}
	// N = max_nums
	fmt.Println("N is", N)
	tmp = N
	res := 0
	for i := 0; tmp > 0; i++ {
		num := tmp % 10
		res += dp[num] * int(math.Pow(float64(dp[9]), float64(i)))
		tmp = tmp / 10
	}
	fmt.Println(dp)
	fmt.Println(res)

	return res

}

func main() {
	atMostNGivenDigitSet([]string{"3", "4", "6", "7", "9"}, 4170)
}
