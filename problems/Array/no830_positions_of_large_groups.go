package arr

import "fmt"

func LargeGroupPositions(S string) [][]int {
	sLen := len(S)
	fmt.Println(sLen)
	ans := [][]int{}
	if sLen < 3 {
		return ans
	}
	cur := S[0]
	start := 0
	end := start
	for i := 0; i <= sLen-2; i++ {
		fmt.Println(i, start, end)
		if S[i] == S[i+1] {
			if cur == S[i] {
				end++
				fmt.Println(start, end)
			} else {
				if end-start+1 >= 3 {
					ans = append(ans, []int{start, end})
				}
				cur = S[i]
				start = i
				end = i
				i--
			}
		}
	}
	if end-start+1 >= 3 {
		ans = append(ans, []int{start, end})
	}
	return ans
}

func LargeGroupPositions2(S string) [][]int {
	sLen := len(S)
	ans := [][]int{}
	i := 1
	count := 0
	for i < sLen {
		if S[i] == S[i-1] {
			count++
		} else {
			if count >= 2 {
				ans = append(ans, []int{i - 1 - count, i - 1})
			}
			count = 0
		}
		i++
	}
	if count >= 2 {
		ans = append(ans, []int{i - 1 - count, i - 1})
	}
	return ans
}
