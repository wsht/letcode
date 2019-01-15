package arr

import (
	"fmt"
	"math"
)

//这个问题需要一个数组来记录该点是否被访问过，
// 由于题目中没有说不可以操作原数组， 而且每个item只访问一次
// 那么我们既可以用一个新数组来记录，也可以用原数组来记录是否访问

//这个方法多做了一个循环，如果一个形成了闭环，那么不应该再有其他额外的数字存在
func ArrayNesting(nums []int) int {
	n := len(nums)
	access := make([]int, n)
	for i := 0; i < n; i++ {
		if access[i] != 0 {
			continue
		}
		fmt.Println(i, nums[i])
		access[i]++
		pre := i
		index := nums[i]
		for access[index] == 0 {
			access[index] = access[pre] + 1
			pre = index
			index = nums[index]
		}
	}
	max := 0
	for i := 0; i < n; i++ {
		max = int(math.Max(float64(access[i]), float64(max)))
	}
	// fmt.Println(access)
	return max
}

//我们同样也可以不使用额外空间，使用nums本身来计算
func ArrayNestingWithoutExtraSpace(nums []int) int {
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			start := nums[i]
			count := 0
			for nums[start] >= 0 {
				tmp := start
				start = nums[start]
				count++
				nums[tmp] = -1
			}
			max = int(math.Max(float64(max), float64(count)))
		}
	}

	return max
}
