package ln

//space O(n) time O(n)
func ReorderListWithSlice(head *ListNode) {
	dummy := &ListNode{}
	record := []*ListNode{}
	for head != nil {
		record = append(record, head)
		head = head.Next
	}
	i := 0
	tmp := dummy
	for ; i < len(record)/2; i++ {
		tmp.Next = record[i]
		tmp = tmp.Next
		tmp.Next = (record[len(record)-i-1])
		tmp = tmp.Next
	}
	if len(record)%2 == 1 {
		tmp.Next = record[i]
		tmp = tmp.Next
	}
	tmp.Next = nil
}

//space O(1) time O(n)
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//1.找到节点中间点
	middle := SplitInMiddle(head)
	//2。reverse list
	newTail := ReverseList(middle)
	tmp := head

	for newTail != nil {
		nextTail := newTail.Next
		newTail.Next = nil
		copyTmpNext := tmp.Next
		tmp.Next = newTail
		newTail.Next = copyTmpNext
		newTail = nextTail
		if tmp.Next.Next == nil {
			tmp = tmp.Next
			break
		} else {
			tmp = tmp.Next.Next
		}
	}
	tmp.Next = newTail
}
