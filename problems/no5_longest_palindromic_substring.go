package main

import (
	"fmt"
	"math"
)

/**
time O(n^3)
spqce O(n)
*/
func longestPalindrome(s string) string {
	strLen := len(s)
	// fmt.Println("strlne", strLen)
	if strLen <= 1 {
		return s
	}

	for readed := 0; readed < strLen; readed++ {
		curStrLen := (strLen - readed)
		// fmt.Println("currstrlen", curStrLen)
		for i := 0; i <= readed; i++ {
			// fmt.Println("start", i, i+curStrLen)
			start := i
			end := i + curStrLen - 1

			if curStrLen <= 1 {
				return s[0:1]
			}

			for j := 0; j < curStrLen/2; j++ {
				if s[start+j] != s[end-j] {
					break
				}
				if j+1 != curStrLen/2 {
					continue
				}
				return s[start : end+1]
			}
		}
	}
	return ""
}

/**
time O(n^2)
space O(n^2)
*/

func longestPalindrome_shffixTree(s string) string {
	str1 := []byte{byte(" "[0])}
	str2 := []byte{byte(" "[0])}
	strlen := len(s)
	for i := 0; i < strlen; i++ {
		str1 = append(str1, s[i])
		str2 = append(str2, s[strlen-i-1])
	}

	return longestCommonSubstr(str1, str2)
}

//todo check if the substring's indices are the same as the reversed substring's original indices
/**
最大公共子串查找
应用次方法 有一个问题
input:abcdasdfghjkldcba
output: dcba
expected:"a"

https://en.wikipedia.org/wiki/Longest_common_substring_problem
*/
func longestCommonSubstr(str1 []byte, str2 []byte) string {

	str1Len := len(str1)
	str2Len := len(str2)
	// ret := make(map[string]int)
	arrayMap := make([][]int, str1Len)
	ret2 := ""
	for i := 0; i < str1Len; i++ {
		arrayMap[i] = make([]int, str2Len)
	}

	z := 0
	for i := 1; i < str1Len; i++ {
		for j := 1; j < str2Len; j++ {
			if str1[i] == str2[j] {
				if i == 1 || j == 1 {
					arrayMap[i][j] = 1
				} else {
					arrayMap[i][j] = arrayMap[i-1][j-1] + 1
				}
				if arrayMap[i][j] > z {
					z = arrayMap[i][j]
					ret2 = string(str1[i-z+1 : i+1])
					// ret[string(str1[i-z+1:i+1])]++
				} else if arrayMap[i][j] == z {
					// ret[string(str1[i-z+1:i+1])]++
					ret2 = string(str1[i-z+1 : i+1])
				}
			} else {
				arrayMap[i][j] = 0
			}
		}
	}
	// fmt.Println(arrayMap)
	for _, value := range arrayMap {
		fmt.Println(value)
	}
	return ret2
}

/**
我们知道，回文是以它的中心点为镜像，因此，回文是可以由其中心点展开的，所以，这里只有2n-1个这样的中心点

time: O(n^2)
*/
func longestPalindrome_4(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))

		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	fmt.Println(start, end)
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

/**
An O(N) Solution (Manacher’s Algorithm):
https://articles.leetcode.com/longest-palindromic-substring-part-ii/
todo understand
*/

func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}

	ret := "^"
	for i := 0; i < n; i++ {
		ret += "#" + string(s[i])
	}
	ret += "#$"
	return ret
}

func longestPalindrome_5(s string) string {
	T := preProcess(s)
	n := len(T)
	p := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		i_mirror := 2*C - i

		if R > i {
			p[i] = int(math.Min(float64(R-i), float64(p[i_mirror])))
		} else {
			p[i] = 0
		}

		//Attempt to expand palidrome centered at i
		for T[i+1+p[i]] == T[i-1-p[i]] {
			p[i]++
		}

		//if palindrome centered at i expand past R
		//adjust center based on expanded palindrome
		if i+p[i] > R {
			C = i
			R = i + p[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i := 0; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}
	fmt.Println(p)
	fmt.Println(centerIndex, maxLen)
	return s[(centerIndex-1-maxLen)/2 : (centerIndex-1+maxLen)/2]
}

// func main() {
// fmt.Println(longestPalindrome("eabcb"))
// fmt.Println(longestPalindrome_4("abacd"))
// fmt.Println(longestPalindrome_5("babcbabcbaccba"))
// }

//todo suffix trees
