package tn

func SumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sum += SumOfLeftLeaves(root.Left)
		}
	}
	sum += SumOfLeftLeaves(root.Right)
	return sum
}
