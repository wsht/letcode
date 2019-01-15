package ln

func ReverseList(head *ListNode) *ListNode {
	result := &ListNode{}
	result.Next = head
	if head == nil {
		return nil
	}
	for head.Next != nil {
		tmp := result.Next
		next := head.Next
		head.Next = next.Next
		next.Next = tmp
		result.Next = next
	}
	return result.Next
}
