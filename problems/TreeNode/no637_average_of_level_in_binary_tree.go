package tn

func AverageOfLevels(root *TreeNode) []float64 {
	record := [][]int{}
	mapTree(root, 0, &record)
	n := len(record)
	ret := make([]float64, n)
	for i := 0; i < len(record); i++ {
		ret[i] = float64(record[i][0]) / float64(record[i][1])
	}

	return ret
}

func mapTree(root *TreeNode, depth int, record *[][]int) {
	if root == nil {
		return
	}
	if len(*record) < depth+1 {
		*record = append(*record, []int{0, 0})
	}
	(*record)[depth][0] += root.Val
	(*record)[depth][1]++

	mapTree(root.Left, depth+1, record)
	mapTree(root.Right, depth+1, record)
}

func AverageOfLevelsBFS(root *TreeNode) []float64 {
	res := []float64{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		sum, count := 0, 0
		tmp := []*TreeNode{}
		for len(queue) != 0 {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			sum += node.Val
			count++
			if node.Left != nil {
				tmp = append(tmp, node)
			}
			if node.Right != nil {
				tmp = append(tmp, node)
			}
		}
		queue = tmp
		res = append(res, float64(sum)/float64(count))
	}
	return res
}
