package main

import (
	"math"
)

/**
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea(height []int) int {
	maxarea := 0
	l := 0
	r := len(height) - 1
	for l < r {
		maxarea = int(math.Max(float64(maxarea), math.Min(float64(height[l]), float64(height[r]))*float64(r-l)))
		//to find bigger value
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxarea
}
