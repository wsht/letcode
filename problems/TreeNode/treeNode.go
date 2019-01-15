package tn

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) Print() {
	fmt.Println(this.Val)
	if this.Left == nil && this.Right == nil {
		return
	}
	fmt.Println("left")
	if this.Left == nil {
		fmt.Println("nil")
	} else {
		this.Left.Print()
	}
	fmt.Println("right")
	if this.Right == nil {
		fmt.Println("nil")
	} else {
		this.Right.Print()
	}
}

func getLeft(si int) int {
	return 2*si + 1
}
func getRight(si int) int {
	return 2*si + 2
}

func InitFromSlice(list []int) *TreeNode {
	treeList := []*TreeNode{}
	n := len(list)
	for i := 0; i < n; i++ {
		node := &TreeNode{}
		if list[i] == 0 {
			node = nil
		} else {
			node.Val = list[i]
		}
		treeList = append(treeList, node)
	}

	for i := 0; i < len(treeList); i++ {
		left := getLeft(i)
		right := getRight(i)
		if left < n {
			treeList[i].Left = treeList[left]
		}
		if right < n {
			treeList[i].Right = treeList[right]
		}
	}

	return treeList[0]
}
