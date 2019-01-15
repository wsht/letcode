package main

import (
	"letcode/problems/ListNode"
	"letcode/problems/TreeNode"
)

func sortedListToBST(head *ln.ListNode) *tn.TreeNode {
	treeHead := &tn.TreeNode{}
	cur := treeHead.Left
	for head != nil {
		val := head.Val
		tmp := &tn.TreeNode{Val: val}
		if cur == nil {
			treeHead.Left = tmp
			cur = treeHead.Left
		} else {
			if cur.Val < val {
				if cur.Left == nil || cur.Right != nil {
					tmp.Left = cur
					cur = tmp
					treeHead.Left = cur
				} else {
					cur.Right = tmp
				}
			}
		}
		head = head.Next
	}
	return treeHead.Left
}

func main() {
	list := ln.InitFromSlice([]int{-10, -3, 0, 5, 9})
	bst := sortedListToBST(list)
	bst.Print()
}
