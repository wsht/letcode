package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	curNode := head
	dummy := &ListNode{}
	preHead := dummy
	dummy.Next = head
	for curNode != nil {
		if curNode.Next == nil {
			break
		}
		nextHead := curNode.Next.Next
		curNode.Next.Next = curNode
		preHead.Next = curNode.Next
		preHead = curNode
		curNode.Next = nextHead
		curNode = nextHead
	}
	return dummy.Next
}

func swapPairs_recursion(head *ListNode) *ListNode {
	list := head
	if head != nil && head.Next != nil {
		list = head.Next
		head.Next = swapPairs_recursion(list.Next)
		list.Next = head
	}
	return list
}

// func main() {
// 	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
// 	// head := &ListNode{Val: 1, Next: nil}

// 	result := swapPairs(head)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }
