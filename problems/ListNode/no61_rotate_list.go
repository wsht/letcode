package ln

import (
	"fmt"
)

func RotateRight(head *ListNode, k int) *ListNode {
	fast := head
	tail := head
	len := 0
	for fast != nil && fast.Next != nil {
		len += 2
		if fast.Next.Next == nil {
			tail = fast.Next
		} else if fast.Next.Next.Next == nil {
			tail = fast.Next.Next
		}
		fast = fast.Next.Next
	}
	if fast != nil {
		len += 1
	}
	if len == 0 {
		return head
	}
	k = k % len
	if k == 0 {
		return head
	}
	fmt.Println(tail)
	//build circle
	tail.Next = head

	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	for i := 0; i < len-k-1; i++ {
		cur = cur.Next
	}

	dummy.Next = cur.Next
	cur.Next = nil
	return dummy.Next
	// cur = dummy.Next

	// for i := 0; i < k-1; i++ {
	// 	cur = cur.Next
	// }
	// cur.Next = head
	// return dummy.Next
}
