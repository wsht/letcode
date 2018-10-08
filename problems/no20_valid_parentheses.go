package main

import (
	"fmt"
)

func checkClose(cur, next string) bool {
	switch cur {
	case "(":
		if next == ")" {
			return true
		}
	case "[":
		if next == "]" {
			return true
		}
	case "{":
		if next == "}" {
			return true
		}
	}

	return false
}

func isValid(s string) bool {
	strLen := len(s)
	if strLen%2 != 0 {
		return false
	}
	start := 0
	end := start + 1
	for start < strLen {
		fmt.Println(start, end)
		tmpStart := start
		for !checkClose(s[tmpStart:end], s[end:end+1]) {
			tmpStart++
			end = tmpStart + 1
			if end == strLen {
				return false
			}
			continue
		}
		for tmpStart > start {
			tmpStart--
			end++
			if end == strLen {
				return false
			}
			if !checkClose(s[tmpStart:tmpStart+1], s[end:end+1]) {
				return false
			}
		}
		start = end + 1
		end = start + 1
	}

	return true
}

func isValid2(s string) bool {
	var mappings = make(map[string]string)
	mappings[")"] = "("
	mappings["}"] = "{"
	mappings["]"] = "["

	stack := []string{}
	for i := 0; i < len(s); i++ {
		tmpS := s[i : i+1]
		openChar, ok := mappings[tmpS]
		if ok {
			topChar := "#"
			if len(stack) > 0 {
				topChar, stack = stack[len(stack)-1 : len(stack)][0], stack[:len(stack)-1]
			}
			if topChar != openChar {
				return false
			}
		} else {
			stack = append(stack, tmpS)
		}
	}
	return len(stack) == 0
}

// func main() {
// 	s := "(([]){})"
// 	fmt.Println("string:", s, " is valid", isValid2(s))
// }
