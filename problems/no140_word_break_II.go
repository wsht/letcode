package main

import (
	"fmt"
)

func wordBreak_Process(s string, wordDict []string, curString string, result *[]string, deep int) {
	// fmt.Println("cur deep", deep)
	if len(s) == 0 {
		// fmt.Println("result append ", curString)
		*result = append(*result, curString)
		// fmt.Println(result)
		return
	}
	deep += 1
	for i := 0; i < len(wordDict); i++ {
		cur_len := len(wordDict[i])
		if len(s) >= cur_len {
			// fmt.Println(s[0:cur_len], deep)
			if s[0:cur_len] == wordDict[i] {
				new_cur_string := ""
				if curString == "" {
					new_cur_string = wordDict[i]
				} else {
					new_cur_string = curString + " " + wordDict[i]
				}

				wordBreak_Process(s[cur_len:], wordDict, new_cur_string, result, deep)
			}
		}
	}
}

func wordBreak(s string, wrodDict []string) []string {
	result := []string{}
	wordBreak_Process(s, wrodDict, "", &result, 0)
	return result
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, wordDict))
}
