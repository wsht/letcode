package main

import (
	"letcode/problems/ListNode"
)

func partition(head *ln.ListNode, x int) *ln.ListNode {
	less, more := &ln.ListNode{}, &ln.ListNode{}
	lessTail, moreTail := &ln.ListNode{}, &ln.ListNode{}
	for head != nil {
		// next := head.Next
		// head.Next = nil
		if head.Val < x {
			if less.Next == nil {
				less.Next = head
			}
			lessTail.Next = head
			lessTail = lessTail.Next
		} else {
			if more.Next == nil {
				more.Next = head
			}
			moreTail.Next = head
			moreTail = moreTail.Next
		}
		// head = next
		head = head.Next
	}
	//防止闭环
	moreTail.Next = nil
	if less.Next == nil {
		return more.Next
	} else {
		lessTail.Next = more.Next
		return less.Next
	}

}

func main() {
	list := ln.InitFromSlice([]int{1, 2, 4, 3, 2, 3})
	list = partition(list, 3)
	list.Print()
}
