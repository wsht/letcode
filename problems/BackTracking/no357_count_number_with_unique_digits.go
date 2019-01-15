package backtrack

func CountNumberWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	total := 0
	visited := make([]bool, 10)
	CountNumberWithUniqueDigitsHelper(&total, n, visited, false, 1)
	return total
}

func CountNumberWithUniqueDigitsHelper(total *int, n int, visited []bool, iszeroStart bool, lens int) {
	if lens > n {
		return
	}

	for i := 0; i <= 9 && lens <= n; i++ {
		if !visited[i] && !iszeroStart {
			visited[i] = true
			*total++
			CountNumberWithUniqueDigitsHelper(total, n, visited, lens == 1 && i == 0, lens+1)
			visited[i] = false
		}
	}
}

//https://leetcode.com/problems/count-numbers-with-unique-digits/discuss/83041/JAVA-DP-O(1)-solution.
func CountNumberWithUniqueDigitsWithDp(n int) int {
	if n == 0 {
		return 1
	}

	res := 10
	uniqueDigits := 9
	availableNumber := 9
	for ; n > 1 && availableNumber > 0; n-- {
		uniqueDigits = uniqueDigits * availableNumber
		res += uniqueDigits
		availableNumber--
	}
	return res
}
