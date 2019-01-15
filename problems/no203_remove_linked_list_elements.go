package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) print() {
	// cur := this
	for this != nil {
		fmt.Println(this.Val)
		this = this.Next
	}

}

func sliceToListNode(list []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for i := 0; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}

	return dummy.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	next := head
	for next != nil {
		if next.Val == val {
			cur.Next = next.Next
			next = cur.Next
		} else {
			cur = next
			next = next.Next
		}
	}
	return dummy.Next
}

func main() {
	result := sliceToListNode([]int{1, 1, 1, 1, 1, 1})
	// result.print()
	result2 := removeElements(result, 1)
	result2.print()
}
