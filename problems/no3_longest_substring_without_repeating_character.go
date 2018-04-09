package main

import (
	"math"
)

/**
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	characterMaps := make(map[byte]int)
	start := 0
	maxLen := 0
	for i := 0; i != len(s); i++ {
		if characterMaps[s[i]] > start {
			start = characterMaps[s[i]]
		}
		characterMaps[s[i]] = i + 1
		maxLen = int(math.Max(float64(maxLen), float64(i-start+1)))
	}

	return maxLen
}

// func main() {
// 	s := "abc"
// 	fmt.Println(lengthOfLongestSubstring(s))
// }
