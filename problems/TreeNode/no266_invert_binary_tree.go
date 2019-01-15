package tn

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

func InvertTreeWithBSF(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		tmp := node.Right
		node.Right = node.Left
		node.Left = tmp
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return root
}
