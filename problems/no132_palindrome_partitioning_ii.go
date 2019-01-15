package main

import (
	"fmt"
	"math"
)

func print_map(s_map [][]int) {
	for i := 0; i < len(s_map); i++ {
		fmt.Println(s_map[i])
	}
}

func minCut(s string) int {
	s_len := len(s)
	s_map := [][]int{}
	for i := 0; i < s_len; i++ {
		s_map = append(s_map, make([]int, s_len))
	}
	for i := 0; i < s_len; i++ {
		for j := s_len - 1; j >= 0; j-- {
			if s[i] == s[j] {
				s_map[i][s_len-1-j] = 1
			}
		}
	}
	print_map(s_map)
	res := 0
	i, j := 0, 0
	for i < s_len {
		fmt.Println(i)
		for s_map[i][j] != 1 {
			j += 1
		}
		if j >= s_len-1 || i >= s_len-1 || s_map[i+1][j+1] != 1 {
			res += 1
			j = 0
			i += 1
		} else {
			i += 1
			j += 1
		}
	}
	return res - 1
}

func minCut2(s string) int {
	return minCut2_helper(s) - 1
}

func minCut2_helper(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len_1 := centerAround(s, i, i)
		len_2 := centerAround(s, i, i+1)
		len := int(math.Max(float64(len_1), float64(len_2)))
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	fmt.Println(start, end)
	if end-start == len(s) {
		return 1
	} else {
		return 1 + minCut2_helper(s[0:start]) + minCut2_helper(s[end+1:len(s)])
	}

}

func centerAround(s string, L int, R int) int {
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

//本题目问题主要是找到最大的回文字串，而不是所有的回文字串

func minCut3(s string) int {
	s_len := len(s)
	s_map := [][]int{}
	for i := 0; i < s_len; i++ {
		s_map = append(s_map, make([]int, s_len))
	}
	//预处理
	for i := 0; i < s_len; i++ {
		for j := s_len - 1; j >= 0; j++ {
			if s[i] == s[j] {
				s_map[i][s_len-1-j] = 1
			}
		}
	}
	print_map(s_map)

	//寻找最大的回文字串
	return 1
}

//好吧 这里我没有解决出来，以上推断是不正确的，暂时放一放

func main() {
	s := "abbbb"
	fmt.Println(minCut2(s))
	// s := "aa"
	// minCut2_helper(s)
}
