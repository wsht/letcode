package main

import (
	"fmt"
	"letcode/problems/ListNode"
)

/*
好吧 需要确定一点的是，这个问题最好的办法是把链表考入到数组中进行排序，尽管他是占空间的，但是在并发环境下，这么做是危险的
*/

func insertionSortList(head *ln.ListNode) *ln.ListNode {
	dummy := &ln.ListNode{}
	dummy.Next = head
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	start := head
	pre := dummy
	start.Next = nil
	for cur != nil {
		if cur.Val < start.Val {
			nextCur := cur.Next
			cur.Next = start
			pre.Next = cur
			cur = nextCur
			pre = dummy
			start = dummy.Next
		} else {
			//最后一个 进行链接
			if start.Next == nil {
				nextCur := cur.Next
				start.Next = cur
				cur.Next = nil
				cur = nextCur
				pre = dummy
				start = dummy.Next
			} else {
				start = start.Next
				pre = pre.Next
			}
		}
	}
	return dummy.Next
}

func insertionSortList(head *ln.ListNode) *ln.ListNode {
	if head == nil {
		return head
	}

	dummy := &ln.ListNode{}
	cur := head
	pre := helper
	next := &ln.ListNode{}
	next = nil
	for cur != nil {
		next = cur.Next
		for pre.Next != nil && pre.Next.val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		pre = dummy
		cur = next
	}
	return dummy.Next
}

func main() {
	list := ln.InitFromSlice([]int{4, 2})
	list = insertionSortList(list)
	fmt.Println("result")
	list.Print()
}
