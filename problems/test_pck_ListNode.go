package main

import (
	"letcode/problems/ListNode"
)

func main() {
	result := ln.InitFromSlice([]int{1, 2, 3})
	result.Print()
	result = ln.AddFront(result, 8)
	result.Print()
}
