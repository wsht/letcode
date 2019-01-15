package tn

import (
	"math"
)

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	ans += 1 + int(math.Max(float64(MaxDepth(root.Left)), float64(MaxDepth(root.Right))))

	return ans
}

type maxDepth struct {
	tn    *TreeNode
	depth int
}

func MaxDepthWithStack(root *TreeNode) int {
	stack := []maxDepth{}
	if root != nil {
		stack = append(stack, maxDepth{tn: root, depth: 1})
	}

	depth := 0
	for len(stack) != 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tn := current.tn
		currDepth := current.depth
		if tn != nil {
			depth = int(math.Max(float64(depth), float64(currDepth)))
			stack = append(stack, maxDepth{tn: tn.Left, depth: depth + 1})
			stack = append(stack, maxDepth{tn: tn.Right, depth: depth + 1})
		}
	}

	return depth
}
