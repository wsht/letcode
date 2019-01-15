package tn

func TrimBST(root *TreeNode, L, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return TrimBST(root.Right, L, R)
	}

	if root.Val > R {
		return TrimBST(root.Left, L, R)
	}

	root.Left = TrimBST(root.Left, L, R)
	root.Right = TrimBST(root.Right, L, R)

	return root
}
