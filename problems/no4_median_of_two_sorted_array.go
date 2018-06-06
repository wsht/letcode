package main

import (
	"fmt"
	"math"
)

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

*/

func findMedianSortedArray(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)
	if m > n { // to ensure m <= n
		tmpnums := nums1
		nums1 = nums2
		nums2 = tmpnums

		tmpLen := m
		m = n
		n = tmpLen
	}
	iMin := 0
	iMax := m
	halfLen := (m + n + 1) / 2

	fmt.Println(iMin, iMax, halfLen)

	for iMin <= iMax {
		i := int(iMin+iMax) / 2
		j := halfLen - i
		fmt.Println("i", i, "j", j)
		if i < iMax && nums2[j-1] > nums1[i] {
			fmt.Println("i is too small")
			iMin = iMin + 1 //i is too small
		} else if i > iMin && nums1[i-1] > nums2[j] {
			fmt.Println("i is too big")
			iMax = iMax - 1 //i is too big
		} else { //i is perfect
			fmt.Println("i is perfet")
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i]
			} else {
				fmt.Println("Max ", nums1[i-1], nums1[j-1])
				maxLeft = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j-1]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return float64(minRight+maxLeft) / 2
		}
	}

	return float64(0)
}

// func main() {
// 	nums1 := []int{3, 4}
// 	nums2 := []int{1}

// 	fmt.Println(findMedianSortedArray(nums1, nums2))
// }
