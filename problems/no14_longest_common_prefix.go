package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	compare := strs[0]
	end := len(compare)
	for i := 0; i < len(compare); i++ {
		if !comparePreFix(strs, i, compare[i]) {
			end = i
			break
		}
	}

	return compare[0:end]
}

func comparePreFix(strs []string, index int, value byte) bool {
	for i := 0; i < len(strs); i++ {
		if index >= len(strs[i]) || value != strs[i][index] {
			return false
		}
	}

	return true
}

/**
less line solution
*/

func longestCommonPrefix2(strs []string) string {
	for idx := 0; len(strs) > 0; idx++ {
		for i := 0; i < len(strs); i++ {
			if idx >= len(strs[i]) || (i > 0 && strs[i][idx] != strs[i-1][idx]) {
				return strs[0][0:idx]
			}
		}
	}
	return ""
}

// func main() {
// 	strs := []string{"aa", "a"}
// 	fmt.Println(longestCommonPrefix2(strs))
// }
