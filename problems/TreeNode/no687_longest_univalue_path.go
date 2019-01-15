package tn

import (
	"math"
)

var total int

func LongestUnivaluePath(root *TreeNode) int {
	sum := 0
	total = 0
	arrowLength(root)
	sum = total
	total = 0
	return sum
}

func arrowLength(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := arrowLength(node.Left)
	right := arrowLength(node.Right)
	arrowLeft, arrowRight := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		arrowLeft += left + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		arrowRight += right + 1
	}
	total = int(math.Max(float64(total), float64(arrowLeft+arrowRight)))
	return int(math.Max(float64(arrowLeft), float64(arrowRight)))
}
