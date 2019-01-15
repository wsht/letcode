package main

import (
	"fmt"
	"letcode/problems/ListNode"
	"math"
)

func splitListToParts(root *ln.ListNode, k int) []*ln.ListNode {
	listLen := 0
	tmp := root
	for tmp != nil {
		listLen++
		tmp = tmp.Next
	}

	every := int(math.Max(float64(1), float64(listLen/k)))

	surplus := int(math.Max(float64(0), float64(listLen-every*k)))
	fmt.Println(every, surplus)

	result := []*ln.ListNode{}
	tmp = root
	for i := 0; i < k; i++ {

		result = append(result, tmp)
		for j := 1; j < every; j++ {
			if tmp != nil {
				tmp = tmp.Next
			}
		}
		if surplus > 0 {
			surplus--
			if tmp != nil {
				tmp = tmp.Next
			}
		}
		if tmp != nil {
			fmt.Println("last tmp val is ", tmp.Val)
			next := tmp.Next
			tmp.Next = nil
			tmp = next
		}
	}

	return result
}

func main() {
	root := ln.InitFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result := splitListToParts(root, 3)
	for i := 0; i < len(result); i++ {
		fmt.Println("index i", i)
		result[i].Print()
	}
}
