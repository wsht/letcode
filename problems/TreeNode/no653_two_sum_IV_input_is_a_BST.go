package tn

func FindTarget(root *TreeNode, k int) bool {
	vmap := map[int]int{}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		_, ok := vmap[k-node.Val]
		//BST 中没有相同的值
		if ok {
			return true
		}
		vmap[node.Val] = 1
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return k == 0
}
