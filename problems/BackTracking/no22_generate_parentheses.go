package backtrack

func GenerateParenthesis(n int) []string {
	ret := []string{}
	GenerateParenthesisHelper(&ret, "", 0, 0, n)
	return ret
}

func GenerateParenthesisHelper(ret *[]string, cur string, open, close, max int) {
	if len(cur) == 2*max {
		*ret = append(*ret, cur)
		return
	}

	if open < max {
		GenerateParenthesisHelper(ret, cur+"(", open+1, close, max)
	}
	if close < open {
		GenerateParenthesisHelper(ret, cur+")", open, close+1, max)
	}
}

// func backtrack(ans *[]string, cur string, open, close, max int) {
// 	if len(cur) == max*2 {
// 		*ans = append(*ans, cur)
// 		return
// 	}
// 	if open < max {
// 		fmt.Println("open < max", cur)
// 		backtrack(ans, cur+"(", open+1, close, max)
// 	}
// 	if close < open {
// 		fmt.Println("close < open", cur)
// 		backtrack(ans, cur+")", open, close+1, max)
// 	}
// }

// func generateParenthesis(n int) []string {
// 	result := []string{}
// 	backtrack(&result, "", 0, 0, n)
// 	return result
// }
