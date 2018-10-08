package main

import (
	"fmt"
)

/*
given a linked list, remove the n-th node from the end of list and return its head.
*/

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func create_list(n int) *ListNode {
	var head ListNode
	tmp := &head
	for i := 1; i <= n; i++ {
		tmp.Val = i
		if i == n {
			tmp.Next = nil
		} else {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}

	return &head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//cal ListNode lenght
	tmp := head
	lenght := 0
	for tmp != nil {
		fmt.Println("not nil")
		lenght++
		tmp = tmp.Next
	}
	tmp = head
	if lenght == n {
		return head.Next
	}
	for i := 0; i < lenght-n-1; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return head
}

func removeNthFromEndWithArray(head *ListNode, n int) *ListNode {
	//n must >= 1 and <= head.len
	cache := make([]*ListNode, n+1)
	tmp := head
	i := 0
	for tmp != nil {
		cache = append(cache, tmp)
		cache = cache[1 : n+1+1]
		i++
		tmp = tmp.Next
	}
	if i == n {
		return head.Next
	}
	cache[0].Next = cache[1].Next
	return head
}

/**
官方版本
用两个指针，其中他们之间的距离是n+1 从而找到消除的前一个点
*/
func removeNthFromEndWithDiscuss(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 1; i <= n+1; i++ {
		first = first.Next //first the len is n + 1
	}
	//next len - n -1
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}

// func main() {
// 	list := create_list(3)
// 	fmt.Println(list)
// 	removeNthFromEndWithArray(list, 1)
// 	fmt.Println(list)
// }
