package main

import (
	"fmt"
)

/**
soluction with O(n^3)
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

func main() {
	fmt.Println(longestPalindrome("eabcb"))
}
