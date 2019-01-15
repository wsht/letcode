package main

import (
	"fmt"
	"letcode/problems/ListNode"
)

func reverseBetween(head *ln.ListNode, m, n int) *ln.ListNode {
	dummy := &ln.ListNode{}
	dummy.Next = head
	if head == nil {
		return dummy.Next
	}

	i := 1
	newHead := dummy
	for head.Next != nil {
		if n == i {
			break
		}
		//start
		if m-1 == i {
			newHead = head
			fmt.Println(newHead)
		}
		if i >= m {
			newHeadNext := newHead.Next
			next := head.Next
			newHead.Next = next
			head.Next = next.Next
			next.Next = newHeadNext
		} else {
			head = head.Next
		}
		i++
	}

	return dummy.Next
}

func main() {
	list := ln.InitFromSlice([]int{1, 2, 3, 4, 5})
	list = reverseBetween(list, 0, 5)
	list.Print()
}
