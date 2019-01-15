package ln

func DeleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	if head == nil || head.Next == nil {
		return head
	}
	pre := dummy
	cur := head
	next := head.Next
	for next != nil {
		if cur.Val == next.Val {
			if next.Next == nil {
				pre.Next = nil
			}
			next = next.Next
			continue
		}
		if cur.Next != next {
			pre.Next = next
			cur = next
			next = next.Next
		} else {
			pre = pre.Next
			cur = cur.Next
			next = next.Next
		}
	}
	return dummy.Next
}
