package main

import (
	"fmt"
	"math"
)

/*
Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the start of the string.
 If there are less than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters,
 then reverse the first k characters and left the other as original.

Example:
Input: s = "abcdefg", k = 2
Output: "bacdfeg"

Restrictions:
The string consists of lower English letters only.
Length of the given string and k will in the range [1, 10000]
*/

//====== version 2
func swap(c []rune, l int, r int) {
	for l < r {
		tmp := c[l]
		c[l] = c[r]
		c[r] = tmp
		l++
		r--
	}
}
func reverseStrII_2(s string, k int) string {
	strlen := len(s)
	start := 0
	strBuf := make([]rune, 0)
	for _, value := range s {
		strBuf = append(strBuf, value)
	}
	for start < strlen {
		end := int(math.Min(float64(start+k-1), float64(strlen-1)))
		fmt.Println(start)
		fmt.Println(end)
		swap(strBuf, start, end)
		start += 2 * k
		fmt.Println(start)
	}

	return string(strBuf)
}

// func main() {
// 	input := "abcdef"
// 	num := 3
// 	// fmt.Println(reverseStrII_1(input, num))

// 	fmt.Println(reverseStrII_2(input, num))
// }

///这个函数对于题目理解错了 所以结果不正确
// func reverseStrII_1(s string, k int) string {
// 	strlen := len(s)
// 	start := 0
// 	// offset := 0
// 	strBuf := make([]rune, 0)
// 	for _, value := range s {
// 		strBuf = append(strBuf, value)
// 	}

// 	for {
// 		surplus := strlen - start
// 		if start >= strlen {
// 			break
// 		}
// 		fmt.Println(surplus)
// 		if surplus < k {
// 			// if surplus > 1 {
// 			fmt.Println("runhere")
// 			loopValue := int(surplus / 2)
// 			for i := 0; i < loopValue; i++ {
// 				tmp := strBuf[start+i]
// 				strBuf[start+i] = strBuf[start+surplus-1-i]
// 				strBuf[start+surplus-1-i] = tmp
// 			}
// 			start += surplus //or break
// 			// }
// 			fmt.Println("always runhere")

// 			// break
// 		} else {
// 			if k > 1 {
// 				index := start - 1
// 				tmp := strBuf[index+k]
// 				strBuf[index+k] = strBuf[index+k-1]
// 				strBuf[index+k-1] = tmp
// 			}

// 			if surplus >= k && surplus < 2*k {
// 				start += surplus //or break
// 				// break
// 			} else {
// 				//转到 item 2 to handle
// 				start += 2 * k
// 			}
// 		}
// 		time.Sleep(100 * time.Millisecond)
// 	}

// 	return string(strBuf)
// }
