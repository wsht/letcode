package main

import (
	"fmt"
	"math"
)

func isInterleave(s1, s2, s3 string) bool {
	result := true
	s1_index, s2_index := 0, 0
	// stack := [][3]int{}

	for i := 0; i < len(s3); i++ {
		if s1_index < len(s1) && s3[i] == s1[s1_index] {
			fmt.Println("check s1", s1_index, s1[s1_index])
			s1_index++
		} else if s2_index < len(s2) && s3[i] == s2[s2_index] {
			fmt.Println("check s2", s2_index, s2[s2_index])
			s2_index++
		} else {
			result = false
			break
		}
	}
	return result
}

//上述方法不能保证恰巧是正确结果，通过增加栈，保存起路径，来解决此问题
func isInterleave2(s1, s2, s3 string) bool {
	result := true
	stack := [][4]int{}
	si := [2]int{0, 0}
	tmp := [4]int{}
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	for i := 0; i < len(s3); i++ {
		if si[0] < len(s1) && s3[i] == s1[si[0]] {
			stack = append(stack, [4]int{0, si[0], si[1], 1})
			si[0]++
		} else if si[1] < len(s2) && s3[i] == s2[si[1]] {
			stack = append(stack, [4]int{1, si[0], si[1], 1})
			si[1]++
		} else {
			for true {
				if len(stack) == 0 {
					return false
				}
				tmp, stack = stack[len(stack)-1], stack[0:len(stack)-1]
				//没有步数可以走，继续pop
				if tmp[3] <= 0 {
					continue
				}
				si = [2]int{tmp[1], tmp[2]}
				i = len(stack)
				s_i := int(math.Abs(float64(1 - tmp[0])))
				if s_i == 0 {
					// fmt.Println(si[s_i], i)
					if si[s_i] < len(s1) && s1[si[s_i]] == s3[i] {
						stack = append(stack, [4]int{0, si[0], si[1], 0})
						si[s_i]++
						break
					}
				}
				if s_i == 1 {
					if si[s_i] < len(s2) && s2[si[s_i]] == s3[i] {
						stack = append(stack, [4]int{1, si[0], si[1], 0})
						si[s_i]++
						break
					}
				}
			}
		}
	}
	// fmt.Println(stack)
	return result
}

func isInterleave_with_dp(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := [][]bool{}
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]bool, len(s2)+1))
	}
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	// for _, item := range dp {
	// 	fmt.Println(item)
	// }
	return dp[len(s1)][len(s2)]
}

func main() {
	// fmt.Println(isInterleave2("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave2("aabd", "abdc", "aabdabcd"))
	fmt.Println(isInterleave_with_dp("aabd", "abdc", "aabdabcd"))
}
