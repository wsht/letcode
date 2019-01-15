package tn

//全局变量
var curNode *TreeNode

func IncreasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	curNode = dummy

	inorder(root)
	return dummy.Right
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	node.Left = nil
	curNode.Right = node
	curNode = node
	inorder(node.Right)
}

func invreasingBST2(root *TreeNode) *TreeNode {
	list := []int{}
	inorder2(root, &list)
	ans := &TreeNode{}
	cur := ans
	for i := 0; i < len(list); i++ {
		cur.Right = &TreeNode{Val: list[i]}
		cur = cur.Right
	}
	return ans
}

func inorder2(node *TreeNode, list *[]int) {
	if node == nil {
		return
	}
	inorder2(node.Left, list)
	*list = append(*list, node.Val)
	inorder2(node.Right, list)
}
