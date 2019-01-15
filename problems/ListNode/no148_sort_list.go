package ln

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//cut the list to two value
	slow, fast := head, head
	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// fmt.Println(pre.Next.Val)
	l2 := SortList(pre.Next)
	pre.Next = nil
	l1 := SortList(head)
	return MergeTwoLists(l1, l2)
}

// func main() {
// 	list := InitFromSlice([]int{4, 2, 1, 3})

// 	list = sortList(list)
// 	fmt.Println("list result")
// 	list.Print()
// }
