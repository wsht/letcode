package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	//init
	count := 0
	state := []int{}
	i, j := 0, 0
	for ; i < len(t); i++ {
		for ; j < len(s); j++ {
			if t[i] == s[j] {
				state = append(state, j)
				j++
				break
			}
		}
	}

	fmt.Println(state)
	if len(state) != len(t) {
		return count
	}
	count++
	endIndex := len(s) - 1
	for len(state) > 0 {
		search := t[len(state)-1]
		startIndex := state[len(state)-1] + 1
		state = state[:len(state)-1]
		fmt.Println(search, " ", startIndex, " ", endIndex)
		for ; startIndex <= endIndex; startIndex++ {
			if search == s[startIndex] {
				fmt.Println(startIndex)
				count++
			}
		}
		endIndex = startIndex - 1
	}

	return count
}

//===============================================================
//to slow
func numDistinct2(s string, t string) int {
	count := 0
	distinctSubsequences2(s, t, &count, 0, 0)
	return count
}

func distinctSubsequences2(s string, t string, count *int, sindex int, tindex int) {
	search := t[tindex]
	for ; sindex < len(s); sindex++ {
		if s[sindex] == search {
			if tindex == len(t)-1 {
				*count++
			} else {
				distinctSubsequences2(s, t, count, sindex+1, tindex+1)
			}
		}
	}
}

//=================================================================
func numDistinct3(s string, t string) int {
	count := 0
	//init
	state := [][]int{}
	for i := 0; i < len(t); i++ {
		state = append(state, []int{})
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				state[i] = append(state[i], j)
			}
		}
		//无此字母，直接返回
		if len(state[i]) == 0 {
			return count
		}
	}
	// fmt.Println(state)
	distinctSubsequences3(state, 0, -1, &count)
	return count
}

func distinctSubsequences3(state [][]int, start int, compare int, count *int) {
	list := state[start]	for _, val := range list {
		if val > compare {
			if start == len(state)-1 {
				*count++
			} else {
				distinctSubsequences3(state, start+1, val, count)
			}
		}
	}
}

func numDistinct4(s string, t string) int {
	var mem = make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		for j := 0; j < len(s)+1; j++ {
			if i == 0 {
				mem[i] = append(mem[i], 1)
			} else {
				mem[i] = append(mem[i], 0)
			}
		}
	}

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				mem[i+1][j+1] = mem[i][j] + mem[i+1][j]
			} else {
				mem[i+1][j+1] = mem[i+1][j]
			}
		}
	}

	fmt.Println(mem)

	return mem[len(t)][len(s)]
}

// func main() {

// 	// fmt.Println("result", numDistinct2("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc"))
// 	fmt.Println("result", numDistinct4("aabb", "ab"))
// }
