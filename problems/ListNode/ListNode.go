package ln

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) Print() {
	for this != nil {
		fmt.Println(this.Val)
		this = this.Next
	}
}

func (this *ListNode) AddFront(val int) {
	head := this
	this.Val = val
	this.Next = head
}

func InitFromSlice(list []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for i := 0; i < len(list); i++ {
		cur.Next = &ListNode{Val: list[i]}
		cur = cur.Next
	}
	return dummy.Next
}

func AddFront(l *ListNode, Val int) *ListNode {
	return &ListNode{Val: Val, Next: l}
}
