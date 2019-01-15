package main

import (
	"fmt"
	"letcode/problems/TreeNode"
	"net/http"
)

func testArray() [][]int {
	arr := [][]int{{4, 5, 6}}
	testHelper(arr)
	arr = append(arr, []int{7, 8, 9})
	return arr
}

func testHelper(arr [][]int) {
	// a := append(arr, []int{1, 2, 3})
}

func main() {
	// fmt.Println(Greedy.Jump([]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}))
	// tree := &tn.TreeNode{}
	// tree.Val = 90
	// tree.Right = &tn.TreeNode{Val: 91}
	// tree.Left = &tn.TreeNode{Val: 68}
	// tmp := tree.Left
	// tmp.Left = &tn.TreeNode{Val: 49}
	// tmp = tmp.Left
	// tmp.Right = &tn.TreeNode{Val: 52}

	// tmp = tree.Left
	// tmp.Right = &tn.TreeNode{Val: 88}
	// tmp = tmp.Right
	// tmp.Left = &tn.TreeNode{Val: 70}
	// fmt.Println(tn.RangeSumBST(tree, 55, 90))
	// obj := tn.Constructor()
	// fmt.Println(obj.Book(10, 20))
	// fmt.Println(obj.Book(10, 20))
	// fmt.Println(obj.Book(10, 20))
	// fmt.Println(obj.Book(10, 20))

	// position := [][]int{{1, 2}, {2, 3}, {6, 1}}
	// fmt.Println(tn.FallingSquares(position))
	tree1 := tn.InitFromSlice([]int{5, 4, 5, 1, 1, 5})
	fmt.Println(tn.LongestUnivaluePath(tree1))

	http.Client()
	http.Serve()
}

// []
