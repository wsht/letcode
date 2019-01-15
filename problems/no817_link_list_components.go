package main

import (
	"fmt"
	"letcode/problems/ListNode"
)

func numComponents(head *ListNode.ListNode, G []int) int {
	cur := head
	GMap := make(map[int]int)
	for i := 0; i < len(G); i++ {
		GMap[G[i]] = 1
	}
	getStart := false
	num := 0
	for cur != nil {
		_, ok := GMap[cur.Val]
		if ok {
			getStart = true
		} else {
			if getStart {
				num++
				getStart = false
			}
		}
		cur = cur.Next
	}
	if getStart == true {
		num++
	}
	return num
}

func numComponents2(head *ListNode.ListNode, G []int) int {
	cur := head
	GMap := make(map[int]int)
	for i := 0; i < len(G); i++ {
		GMap[G[i]] = 1
	}
	getStart := false
	num := 0
	for cur != nil {
		_, ok := GMap[cur.Val]
		if ok {
			if !getStart {
				num++
				getStart = true
			}
		} else {
			getStart = false
		}
		cur = cur.Next
	}

	return num
}

func main() {
	G := []int{4, 3, 1, 5, 2, 7}
	listNode := ListNode.InitFromSlice(G)
	fmt.Println(numComponents(listNode, []int{2, 7, 4, 3}))
	fmt.Println(numComponents2(listNode, []int{2, 7, 4, 3}))
}
