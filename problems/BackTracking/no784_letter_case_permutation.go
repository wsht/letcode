package backtrack

import (
	"strings"
)

func LetterCasePermutation(S string) []string {
	ret := []string{}
	letterCaseHelper(&ret, S, 0)

	return ret
}

func letterCaseHelper(ret *[]string, src string, start int) {

	if start == len(src) {
		*ret = append(*ret, src)
		return
	}

	if src[start] >= '0' && src[start] <= '9' {
		letterCaseHelper(ret, src, start+1)
		return
	}

	tmp := src[start : start+1]
	tmp = strings.ToUpper(tmp)
	src = src[:start] + tmp + src[start+1:]
	letterCaseHelper(ret, src, start+1)

	tmp = strings.ToLower(tmp)
	src = src[:start] + tmp + src[start+1:]
	letterCaseHelper(ret, src, start+1)

}
