package tn

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	if q != nil && p != nil {
		return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}

	if q == nil && p == nil {
		return true
	}

	return false
}
