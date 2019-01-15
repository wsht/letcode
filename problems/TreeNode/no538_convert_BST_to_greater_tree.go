package tn

var total int

func ConvertBST(root *TreeNode) *TreeNode {
	convertBSTHelper(root)
	//notice 这里是全局变量，需要重置为0 否则在循环调用情况下，会累加
	total = 0
	return root
}
func convertBSTHelper(root *TreeNode) {
	if root == nil {
		return
	}
	convertBSTHelper(root.Right)
	root.Val += total
	total = root.Val
	convertBSTHelper(root.Left)
}
