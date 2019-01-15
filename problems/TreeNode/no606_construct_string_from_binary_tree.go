package tn

import (
	"fmt"
	"strconv"
)

func Tree2str2(t *TreeNode) string {
	str := ""
	end := ""
	stack := []*TreeNode{}
	if t == nil {
		return ""
	}
	stack = append(stack, t)
	for len(stack) != 0 {
		node := stack[len(stack)-1]

		stack = stack[:len(stack)-1]
		if node == nil {
			str += "()"
			continue
		}
		str += "(" + strconv.Itoa(node.Val)
		if node.Left != nil || node.Right != nil {
			end += ")"
		} else {
			if len(end) > 0 {
				str += ")" + end[0:1]
				end = end[1:]
			}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		} else if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	str += end
	return str[1 : len(str)-1]
}

func Tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	stack := []*TreeNode{}
	visited := map[*TreeNode]int{}
	str := ""
	stack = append(stack, t)
	for len(stack) != 0 {
		t := stack[len(stack)-1]
		if t == nil {
			str += "()"
			stack = stack[:len(stack)-1]
			continue
		}
		_, ok := visited[t]
		if ok {
			stack = stack[:len(stack)-1]
			str += ")"
		} else {
			visited[t] = 1
			str += "(" + strconv.Itoa(t.Val)
			if t.Left == nil && t.Right == nil {
				str += "()"
			}
			if t.Right != nil {
				stack = append(stack, t.Right)
				stack = append(stack, t.Left)
			} else if t.Left != nil {
				stack = append(stack, t.Left)
			}
		}
	}
	fmt.Println(str)
	ret := ""

	for i := 1; i < len(str)-1; i++ {
		if str[i] == '(' && str[i+1] == ')' && str[i+2] == ')' {
			i += 2
		}
		// fmt.Println(str[i : i+1])
		ret += str[i : i+1]
	}
	return ret
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	res := ""
	res += strconv.Itoa(t.Val)
	if t.Left != nil {
		res += "(" + tree2str(t.Left) + ")"
	} else if t.Left == nil && t.Right != nil {
		res += "()"
	}

	if t.Right != nil {
		res += "(" + tree2str(t.Right) + ")"
	}

	return res
}
