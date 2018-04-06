package main

/**
Write a function that takes a string as input and returns the string reversed.

Example:
Given s = "hello", return "olleh".
*/

func Swap(c []rune, l int, r int) {
	for l < r {
		tmp := c[l]
		c[l] = c[r]
		c[r] = tmp
		l++
		r--
	}
}

func reverseString_I(s string) string {
	strlen := len(s)
	strByte := make([]rune, 0)
	for _, val := range s {
		strByte = append(strByte, val)
	}

	Swap(strByte, 0, strlen-1)

	return string(strByte)
}
