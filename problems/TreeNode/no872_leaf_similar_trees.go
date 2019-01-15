package tn

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1List := []int{}
	r2List := []int{}

	getTreeLeaf(&r1List, root1)
	getTreeLeaf(&r2List, root2)

	if len(r1List) != len(r2List) {
		return false
	}

	for i := 0; i < len(r1List); i++ {
		if r1List[i] != r2List[i] {
			return false
		}
	}

	return true
}

func getTreeLeaf(list *[]int, t *TreeNode) {
	if t == nil {
		return
	}

	if t.Left == nil && t.Right == nil {
		*list = append(*list, t.Val)
		return
	}

	getTreeLeaf(list, t.Left)
	getTreeLeaf(list, t.Right)

}
