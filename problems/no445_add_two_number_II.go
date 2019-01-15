package main

import "letcode/problems/ListNode"

/**
l1 1-2-3-4
l2 7-8-0

first loop
res 4-11-9-1

second loop
res 2-0-1-4
*/
func addTwoNumbers(l1 *ln.ListNode, l2 *ln.ListNode) *ln.ListNode {
	carry, l1Len, l2Len := 0, 0, 0
	p1, p2 := l1, l2
	res := &ListNode.ListNode{}
	res = nil
	for p1 != nil {
		l1Len++
		p1 = p1.Next
	}
	for p2 != nil {
		l2Len++
		p2 = p2.Next
	}

	p1, p2 = l1, l2
	for l1Len > 0 && l2Len > 0 {
		sum := 0
		if l1Len >= l2Len {
			sum += p1.Val
			p1 = p1.Next
			l1Len--
		}
		if l1Len < l2Len {
			sum += p2.Val
			p2 = p2.Next
			l2Len--
		}
		res = addFront(sum, res)
	}

	p1, res = res, nil
	for p1 != nil {
		p1.Val += carry
		res = addFront(p1.Val%10, res)
		carry = p1.Val / 10
		p1 = p1.Next
	}
	if carry > 0 {
		res = addFront(carry, res)
	}

	return res
}

func addFront(val int, l *ln.ListNode) *ln.ListNode {

	return &ln.ListNode{Val: val, Next: l}
}

//同样 我们可以用 stack 解决次问题，这里不在展开
func addTwoNumbers2() {}

func main() {

}
