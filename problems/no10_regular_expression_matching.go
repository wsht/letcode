package main

import "fmt"

func isMatch(s string, p string) bool {
	//>=2
	if len(p) == 0 {
		return len(s) == 0
	}

	first_match := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	} else {
		return first_match && isMatch(s[1:], p[1:])
	}
}

type Reqular_Result int

const (
	NULL Reqular_Result = iota
	TRUE
	FALSE
)

func isMatch_withDp_Top_To_Down(s string, p string) bool {
	memo := [][]Reqular_Result{}
	fmt.Println(memo)
	for i := 0; i < len(s)+1; i++ {
		memo = append(memo, make([]Reqular_Result, len(p)+1))
	}

	return dp(0, 0, s, p, memo)
}

func dp(i int, j int, s string, p string, memo [][]Reqular_Result) bool {
	// memo[0][0] = false
	if memo[i][j] != NULL {
		return memo[i][j] == TRUE
	}
	ans := false
	if j == len(p) {
		ans = i == len(s)
	} else {
		first_match := i < len(s) && (p[j] == s[i] || p[j] == '.')
		if j+1 == len(p) && p[j+1] == '*' {
			ans = (dp(i, j+2, s, p, memo)) || (first_match && dp(i+1, j, s, p, memo))
		} else {
			ans = first_match && dp(i+1, j+1, s, p, memo)
		}
	}
	if ans {
		memo[i][j] = TRUE
	} else {
		memo[i][j] = FALSE
	}

	return ans
}

func isMatch_withDp_BOttom_Up(s, p string) bool {
	dp := [][]bool{}
	for i := 0; i < len(s)+1; i++ {
		dp = append(dp, make([]bool, len(p)+1))
	}

	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p); j >= 0; j-- {
			first_match := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j+1 < len(p) && first_match[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (first_match && dp[i+1][j])
			} else {
				dp[i][j] = first_match && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func main() {
	// fmt.Println(isMatch("aa", ".*"))
	// dp(1, 1, "1", "1", memo)
	// isMatch_withDp("1", "1")
}
