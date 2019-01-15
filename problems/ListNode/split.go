package ln

func SplitInMiddle(head *ListNode) *ListNode {
	pre := head
	fast, slow := pre, pre
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return slow
}
