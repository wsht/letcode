package main

import (
	"fmt"
)

//重复次数
//首先 我们需要确定的就 s2字符串中的字符，肯定在s1中
//其次，我们需要计算出，s1重复多少个周期，可以完全包含s2，
//然后，根据得到的周期关系，算出，最大整数M 使得 s2 * n2 * M 恰好被 s1 * n1所包含

//一个简单的例子，s1 abc 8, s2 aabbcc 1
// func getMaxRepetitons(s1 string, n1 int, s2 string, n2 int) int {
// 	dp := [][]int{}
// 	dp = append(dp, make([]int, len(s1)))
// 	time := 0
// 	s2_start := 0

// 	for j := 0; j < len(s2); j++ {
// 		is_match := false
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] == s2[j] {
// 				is_match = true
// 				j++
// 				dp[time][i] = 1
// 			}
// 		}
// 		if !is_match {
// 			fmt.Println("un match")
// 			return 0
// 		}
// 		time++
// 		dp := append(dp, make([]int, len(s1)))
// 	}
// 	return 1
// }

//暴力解决问题
//time limit out error 在大数据集的时候，这种方法显得比较愚蠢
func getMaxRepetitons2(s1 string, n1 int, s2 string, n2 int) int {
	M := 0
	s1_len, s2_len := len(s1), len(s2)
	i, j := 0, 0
	if n1 == 0 {
		return 0
	}
	for i < s1_len*n1 {
		for j < s2_len*n2 {
			fmt.Println(i%s1_len, j%s2_len)
			if s1[i%s1_len] == s2[j%s2_len] {
				j++
				i++
			} else {
				i++
			}
			if i >= s1_len*n1 {
				if j == s2_len*n2 {
					return M + 1
				}
				return M
			}
		}
		fmt.Println("once loop")
		j = 0
		// i++
		M++
	}
	return M + 1
}

//让我们回到第一种的解决思路，那就是找到他们最原始的关系，从而快速的解决问题
func getMaxRepetitons3(s1 string, n1 int, s2 string, n2 int) int {
	i, j := 0, 0
	dp := [][]int{}
	time := 0
	n := 1
	dp = append(dp, make([]int, len(s1)))
	for j < len(s2) {
		is_match := false
		for i < len(s1) {
			// fmt.Println(i, j)
			dp[time][i] = n
			if s1[i] == s2[j] {
				is_match = true
				i++
				j++
			} else {
				i++
			}
			if j > len(s2)-1 {
				break
			}
		}

		if !is_match || time >= len(s2)-1 {
			return 0
		}
		// fmt.Println(j)
		if j > len(s2)-1 {
			break
		}
		time++
		i = 0
		dp = append(dp, make([]int, len(s1)))
	}
	// n_tmp := n
	match_count := 0
	for ; i < len(s1); i++ {
		is_match := false
		for j = 0; j < len(s2); j++ {
			if i > len(s1)-1 {
				break
			}
			// fmt.Println("new", i, j)
			// fmt.Println(match_count)
			if s1[i] == s2[j] {
				is_match = true
				i++
				match_count++
			}
		}
		if !is_match {
			break
		}
		i--
	}
	// fmt.Println(match_count)
	n += match_count / len(s2)
	is_just := 0
	if match_count%len(s2) > 0 {
		is_just = 1
	}

	return (n1*n/(time+1) + is_just) / n2
}

func main() {
	fmt.Println(getMaxRepetitons3("baba", 11, "baab", 1))

}

// "acb"
// 4
// "ab"
// 2
"baba"
11
"baab"
1
