package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists_compare_one_by_one(lists []*ListNode) *ListNode {
	//check
	alive := len(lists)
	result := &ListNode{Val: 0}
	var tmpIndex int = -1
	var pointer *ListNode = nil
	fmt.Println(alive)
	for alive > 0 {
		for index, list := range lists {
			// fmt.Println(list)
			if list == nil {
				continue
			}

			if tmpIndex == -1 {
				tmpIndex = index
				continue
			}
			if list.Val < lists[tmpIndex].Val {
				tmpIndex = index
			}
		}
		fmt.Println(tmpIndex)
		if tmpIndex == -1 {
			return result.Next
		}
		if pointer == nil {
			result.Next = &ListNode{}
			pointer = result.Next
			pointer.Val = lists[tmpIndex].Val
		} else {
			pointer.Next = &ListNode{}
			pointer = pointer.Next
		}

		pointer.Val = lists[tmpIndex].Val

		lists[tmpIndex] = lists[tmpIndex].Next
		if lists[tmpIndex] == nil {
			tmpIndex = -1
			alive--
		}
	}

	return result.Next
}

func mergeKLists_compare_one_by_one_2(lists []*ListNode) *ListNode {
	alive := len(lists)
	result := &ListNode{Val: 0}
	var tmpIndex int = -1
	var pointer *ListNode = nil

	for alive > 0 {
		for index, list := range lists {
			// fmt.Println(list)
			if list == nil {
				continue
			}

			if tmpIndex == -1 {
				tmpIndex = index
				continue
			}
			if list.Val < lists[tmpIndex].Val {
				tmpIndex = index
			}
		}

		if tmpIndex == -1 {
			return result.Next
		}
		if pointer == nil {
			result.Next = lists[tmpIndex]
			pointer = result.Next
		} else {
			pointer.Next = lists[tmpIndex]
			pointer = pointer.Next
		}

		lists[tmpIndex] = lists[tmpIndex].Next
		if lists[tmpIndex] == nil {
			tmpIndex = -1
			alive--
		}
	}

	return result.Next
}

//可以是用PriorityQueue 进行排序，降K个list head压入队列中，进行排序
//O(nlog(k))

// func main() {
// 	// lists := []*ListNode{
// 	// 	&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
// 	// 	&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
// 	// 	&ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
// 	// }
// 	lists := []*ListNode{nil, &ListNode{Val: 1, Next: nil}, nil}
// 	test := mergeKLists_compare_one_by_one_2(lists)
// 	for test != nil {
// 		fmt.Println(test.Val)
// 		test = test.Next
// 	}

// }
