package ln

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := &ListNode{}
	result = cur
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		// cur.Next = l2.Next
		cur.Next = l2
	}
	return result.Next
}

//This solution is not a tail-recursive, the stack will overflow while the list is too long :)
func MergeTwoLists_recursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists_recursive(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoLists_recursive(l2.Next, l1)
		return l2
	}
}

// func main() {
// 	result := mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}})
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}

// }
