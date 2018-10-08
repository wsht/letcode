package main

import (
	"fmt"
)

func backtrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		fmt.Println("open < max", cur)
		backtrack(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		fmt.Println("close < open", cur)
		backtrack(ans, cur+")", open, close+1, max)
	}
}

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)
	return result
}

// func main() {
// 	fmt.Println("result:", generateParenthesis(2))

// }
