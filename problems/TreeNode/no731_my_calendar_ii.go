package tn

import (
	"math"
)

type MyCalendarTwo struct {
	start   int
	end     int
	overlap bool
	left    *MyCalendarTwo
	right   *MyCalendarTwo
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if !insertable(this, start, end) {
		return false
	}
	this = insert(this, start, end)
	return true
}

func insert(node *MyCalendarTwo, start, end int) *MyCalendarTwo {
	if start >= end {
		return node
	}

	if node == nil {
		return &MyCalendarTwo{start: start, end: end}
	}

	if start >= node.end {
		node.right = insert(node.right, start, end)
		return node
	} else if end <= node.start {
		node.left = insert(node.left, start, end)
		return node
	} else {
		node.overlap = true
		a := int(math.Min(float64(node.start), float64(start)))
		b := int(math.Max(float64(node.start), float64(start)))
		c := int(math.Min(float64(node.end), float64(end)))
		d := int(math.Max(float64(node.end), float64(end)))
		node.left = insert(node.left, a, b)
		node.right = insert(node.right, c, d)
		node.start = b
		node.end = c
		return node
	}
	// return node
}

func insertable(node *MyCalendarTwo, start, end int) bool {
	if start >= end {
		return true
	}
	if node == nil {
		return true
	}
	if start >= node.end { //check right side
		return insertable(node.right, start, end)
	} else if end <= node.start {
		return insertable(node.left, start, end)
	} else {
		if node.overlap {
			return false
		}
		if start >= node.start && end <= node.end {
			return true
		}

		return insertable(node.left, start, node.start) && insertable(node.right, node.end, end)
	}
}
