package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	if head == nil || head.Next == nil {
		return head
	}
	cur, next := head, head.Next
	for next != nil {
		if cur.Val == next.Val {
			if next.Next == nil {
				cur.Next = next.Next
			}
			next = next.Next

		} else {
			fmt.Println(cur.Val, next.Val)
			cur.Next = next
			cur = cur.Next
			next = next.Next
		}
	}
	// if next.Val == cur.Val {
	// 	cur.Next = next.Next
	// }
	return dummy.Next
}

func main() {
	result := deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}})
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
