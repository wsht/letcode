package Greedy

//最小跳跃次数可以到达数组的最后一个位置

func Jump(nums []int) int {

	step := 1
	maxIndex := nums[0]
	// minIndex := 0
	nextMaxIndex := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if i+nums[i] > nextMaxIndex {
			nextMaxIndex = i + nums[i]
		}
		if i == maxIndex {
			// minIndex = maxIndex
			maxIndex = nextMaxIndex
			step++
		}
		// fmt.Println(i, maxIndex, nextMaxIndex, step)
		//这里相当于一个小优化，防止进行无效的循环
		if nextMaxIndex >= n-1 {
			if maxIndex < nextMaxIndex {
				return step + 1
			} else {
				return step
			}
		}
	}
	return 0
}
