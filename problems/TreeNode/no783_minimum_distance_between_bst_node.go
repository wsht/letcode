package tn

import (
	"fmt"
	"math"
)

func MinDiffInBST(root *TreeNode) int {
	min := math.MaxInt32
	pre := -1
	MinDiffInBSTHelper2(root, &min, &pre)
	// MinDiffInBSTHelper(root, &min)
	return min
}

//找到最小的差值点，根据二叉查找树的特性，他的左子点是是比根节点小的，
// 右子点，是比他大的点。
//那么我们可以知道，他的左子点的最有节点，以及右节点的最左节点是和根节点最相近的值
func MinDiffInBSTHelper(root *TreeNode, min *int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		tmp := root.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		*min = int(math.Min(float64(*min), float64(tmp.Val-root.Val)))
	}

	if root.Left != nil {
		tmp := root.Left
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		*min = int(math.Min(float64(*min), float64(root.Val-tmp.Val)))
	}

	MinDiffInBSTHelper(root.Left, min)
	MinDiffInBSTHelper(root.Right, min)
}

//还是没有明白他的遍历过程
func MinDiffInBSTHelper2(root *TreeNode, ret, pre *int) {
	if root == nil {
		return
	}

	//左 的右子串
	MinDiffInBSTHelper2(root.Left, ret, pre)
	fmt.Println(*ret, *pre, root.Val)

	if *pre >= 0 {
		if *ret > root.Val-*pre {
			*ret = root.Val - *pre
		}
	}
	*pre = root.Val

	MinDiffInBSTHelper2(root.Right, ret, pre)

}

// tree := &tn.TreeNode{}
// 	tree.Val = 90
// 	tree.Right = &tn.TreeNode{Val: 91}
// 	tree.Left = &tn.TreeNode{Val: 68}
// 	tmp := tree.Left
// 	tmp.Left = &tn.TreeNode{Val: 49}
// 	tmp = tmp.Left
// 	tmp.Right = &tn.TreeNode{Val: 52}

// 	tmp = tree.Left
// 	tmp.Right = &tn.TreeNode{Val: 88}
// 	tmp = tmp.Right
// 	tmp.Left = &tn.TreeNode{Val: 70}
// 	fmt.Println(tn.MinDiffInBST(tree))
