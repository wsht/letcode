package tn

func RangeSumBST(root *TreeNode, L, R int) int {
	ret := 0
	RangeSumBSTHelper(root, L, R, &ret)
	return ret
}

func RangeSumBSTHelper(root *TreeNode, L, R int, ret *int) {
	if root == nil {
		return
	}
	if root.Val > R {
		RangeSumBSTHelper(root.Left, L, R, ret)
	}

	if root.Val < L {
		RangeSumBSTHelper(root.Right, L, R, ret)
	}

	if root.Val >= L && root.Val <= R {
		*ret += root.Val
		//小优化，防止无用的查找
		if root.Val != L && root.Val != R {
			RangeSumBSTHelper(root.Left, L, R, ret)
			RangeSumBSTHelper(root.Right, L, R, ret)
		} else if root.Val == L {
			RangeSumBSTHelper(root.Right, L, R, ret)
		} else if root.Val == R {
			RangeSumBSTHelper(root.Left, L, R, ret)
		}

	}
}

func RangeSumBSTWithStack(root *TreeNode, L, R int) int {
	ans := 0
	stack := []*TreeNode{}
	stack = append(stack, root)
	node := &TreeNode{}
	node = nil
	for len(stack) != 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if L <= node.Val && node.Val <= R {
			ans += node.Val
		}
		if L < node.Val {
			stack = append(stack, node.Left)
		}

		if node.Val < R {
			stack = append(stack, node.Right)
		}
	}

	return ans
}
