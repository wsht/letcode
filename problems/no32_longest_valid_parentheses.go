package main

import (
	"math"
)

func longestValidParentheses(s string) int {

	stack := []string{}
	total := 0
	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		//open add to stack
		if str == "(" {
			stack = append(stack, str)
		} else { //close pop stack
			if len(stack) == 0 {
				continue
			} else {
				stack = stack[:len(stack)-1]
				total += 2
			}
		}
	}

	return total
}

/**
记录每个状态
*/
func longestValidParentheses2(s string) int {
	stack := []int{}
	state := []int{}
	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		state = append(state, 0)
		if str == "(" {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				continue
			} else {
				tmpIndex := -1
				tmpIndex, stack = stack[len(stack)-1 : len(stack)][0], stack[:len(stack)-1]
				state[tmpIndex] = 1
				state[i] = 1
			}
		}
	}
	i := 0
	result := []int{0}
	for _, val := range state {
		if val == 0 {
			i++
			result = append(result, 0)
			continue
		}
		result[i] += val
	}
	total := 0
	for _, val := range result {
		if total < val {
			total = val
		}
	}
	return total
}

//对于方法二的改良版本，不在需要state 数组
func longestValidParentheses3(s string) int {
	maxlen := 0
	stack := []int{}
	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		if str == "(" {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				tmpIndex := 0
				tmpIndex, stack = stack[len(stack)-1 : len(stack)][0], stack[:len(stack)-1]
				maxlen = int(math.Max(float64(maxlen), float64(i-tmpIndex)))
			}
		}
	}
	return maxlen
}

//动态规划方法
func longestValidParentheses4(s string) int {
	maxlen := 0
	var dp = make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2]
				} else {
					dp[i] = 0
				}
				dp[i] += 2
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxlen = int(math.Max(float64(maxlen), float64(dp[i])))
		}
	}
	return maxlen
}

//time O(n), space O(1)
func longestValidParentheses5(s string) int {
	left, right, maxlen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxlen = int(math.Max(float64(maxlen), float64(2*right)))
		} else if right >= left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxlen = int(math.Max(float64(maxlen), float64(2*right)))
		} else if left >= right {
			left, right = 0, 0
		}
	}
	return maxlen
}

// func main() {
// 	s := "()(()"
// 	fmt.Println(longestValidParentheses2(s))
// }
