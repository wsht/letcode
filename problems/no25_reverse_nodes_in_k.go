package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func result_dump(tmp_tail *ListNode, k int) {
	fmt.Print("return tail result ", k, "->")
	for tmp_tail != nil {
		fmt.Print(tmp_tail.Val, ",")
		tmp_tail = tmp_tail.Next
	}
	fmt.Print("\n")
}

func reverseKGroup(head *ListNode, k int, total_k int) *ListNode {
	tail := head
	// if k <= 0 {
	// 	return nil
	// }
	fmt.Println(head, k, total_k)
	if k == 0 || head == nil {
		result_dump(tail, k)
		return nil
	}
	if k == 1 {
		result_dump(tail, k)
		return head
	}

	if k <= 0 && head != nil {
		result_dump(tail, k)
		return head
	}

	for i := 0; i < k-1; i++ {
		tail = tail.Next
		if tail == nil {
			result_dump(tail, k)
			return head
		}
	}
	//todo swap
	var next_tail *ListNode = nil
	if tail != nil {
		next_tail = tail.Next
	}

	next_head := head.Next
	mid_list := reverseKGroup(next_head, k-2, total_k)
	if mid_list == nil {
		tail.Next = head
		if total_k == k {
			head.Next = reverseKGroup(next_tail, k, total_k)
		} else {
			head.Next = nil
		}
	} else {
		tail.Next = mid_list
		for mid_list.Next != nil {
			mid_list = mid_list.Next
		}
		mid_list.Next = head
		if total_k == k {
			head.Next = reverseKGroup(next_tail, k, total_k)
		} else {
			head.Next = nil
		}
	}
	result_dump(tail, k)
	return tail
}

func reverseKGroup_2(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	tail := dummy
	start := prev.Next
	for {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		for i := 0; i < k-1; i++ {
			then := start.Next
			//start向后移动一位 //then 插入首位置
			start.Next = then.Next
			then.Next = prev.Next
			prev.Next = then
		}
		prev = start
		tail = start
		start = prev.Next
	}
}

// func main() {
// 	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
// 	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
// 	// head := &ListNode{Val: 1, Next: nil}

// 	result := reverseKGroup_2(head, 5)
// 	for result != nil {

// 		fmt.Println(result.Val)

// 		result = result.Next
// 	}
// }
