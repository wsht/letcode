package tn

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = MergeTrees(t1.Left, t2.Left)
	t1.Right = MergeTrees(t1.Right, t2.Right)

	return t1
}

/**
In the current approach, we again traverse the two trees, but this time we make use of a stackstack to do so instead of making use of recursion. Each entry in the stackstack strores data in the form [node_{tree1}, node_{tree2}][node
tree1
​
 ,node
tree2
​
 ]. Here, node_{tree1}node
tree1
​
  and node_{tree2}node
tree2
​
  are the nodes of the first tree and the second tree respectively.

We start off by pushing the root nodes of both the trees onto the stackstack. Then, at every step, we remove a node pair from the top of the stack. For every node pair removed, we add the values corresponding to the two nodes and update the value of the corresponding node in the first tree. Then, if the left child of the first tree exists, we push the left child(pair) of both the trees onto the stack. If the left child of the first tree doesn't exist, we append the left child(subtree) of the second tree to the current node of the first tree. We do the same for the right child pair as well.

If, at any step, both the current nodes are null, we continue with popping the next nodes from the stackstack.


Complexity Analysis

Time complexity : O(n)O(n). We traverse over a total of nn nodes. Here, nn refers to the smaller of the number of nodes in the two trees.

Space complexity : O(n)O(n). The depth of stack can grow upto nn in case of a skewed tree.
*/
func MergeTreesWithStack(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	stack := [][]*TreeNode{}
	stack = append(stack, []*TreeNode{t1, t2})
	for len(stack) != 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if t[0] == nil || t1 == nil {
			continue
		}

		t[0].Val += t[1].Val
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			stack = append(stack, []*TreeNode{t[0].Left, t[1].Left})
		}
		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			stack = append(stack, []*TreeNode{t[0].Right, t[1].Right})
		}
	}

	return t1
}
