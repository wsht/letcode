package tn

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if val > root.Val {
		return SearchBST(root.Right, val)
	} else {
		return SearchBST(root.Left, val)
	}
}
