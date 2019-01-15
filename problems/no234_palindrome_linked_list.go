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

func isPalindrome(head *ListNode) bool {
	len := 0
	fast, low := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		len += 2
	}
	if fast != nil && fast.Next == nil {
		len += 1
	}
	// fmt.Println(len)
	if len == 0 {
		return true
	}
	if len == 1 {
		return true
	}
	if len%2 == 1 {
		low = low.Next
	}

	reverse_low := &ListNode{}
	reverse_low.Next = low
	for low.Next != nil {
		tmp := reverse_low.Next
		next := low.Next
		low.Next = next.Next
		next.Next = tmp
		reverse_low.Next = next
	}
	reverse_low = reverse_low.Next

	for i := 0; i < len/2; i++ {
		if head.Val != reverse_low.Val {
			return false
		}
		head = head.Next
		reverse_low = reverse_low.Next
	}
	return true
}

func main() {
	result := sliceToListNode([]int{1, 1})
	fmt.Println(isPalindrome(result))

}
