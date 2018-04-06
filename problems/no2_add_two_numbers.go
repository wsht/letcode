package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carray := 0
	sumListNode := &ListNode{}
	currentNode := sumListNode
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carray

		carray = sum / 10
		val := sum % 10

		currentNode.Next = &ListNode{Val: val}
		currentNode = currentNode.Next
	}

	if carray != 0 {
		currentNode.Next = &ListNode{Val: carray}
	}

	return sumListNode.Next
}

// func reverseListNode() {

// }

// func main() {
// 	l1Head := &ListNode{}
// 	d1 := l1Head
// 	l2Head := &ListNode{}
// 	d2 := l2Head
// 	// l2Head := &ListNode{}
// 	// l2 := ListNode{}
// 	// l1tmp := l1
// 	// l2tmp := l2
// 	for i := 1; i <= 3; i++ {
// 		d1.Val = i
// 		d1.Next = &ListNode{}
// 		d1 = d1.Next

// 		d2.Val = 3 * i
// 		d2.Next = &ListNode{}
// 		d2 = d2.Next
// 	}
// 	fmt.Printf("before run l1 address is %v\n", *l1Head)
// 	fmt.Printf("before run l2 address is %v\n", *l2Head)

// 	result := addTwoNumbers(l1Head, l2Head)
// 	fmt.Printf("after run l1 address is %v\n", *l1Head)
// 	fmt.Printf("after run l2 address is %v\n", *l2Head)
// 	for result != nil {
// 		fmt.Println(result.Val)
// 		result = result.Next
// 	}
// }
