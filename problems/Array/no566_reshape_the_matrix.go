package arr

func MatrixReshape(nums [][]int, r, c int) [][]int {
	result := [][]int{}
	nr := len(nums)
	nc := len(nums[0])
	if nc*nr != r*c {
		return nums
	}

	for i := 0; i < r; i++ {
		result = append(result, make([]int, c))
		for j := 0; j < c; j++ {
			num := i*c + j
			result[i][j] = nums[num/nc][num%nc]
		}
	}

	return result
}

func MatrixReshapeWithQueue(nums [][]int, r, c int) [][]int {
	result := [][]int{}
	nr := len(nums)
	nc := len(nums[0])
	if nc*nr != r*c {
		return nums
	}
	queue := []int{}
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			queue = append(queue, nums[i][j])
		}
	}

	for i := 0; i < r; i++ {
		result = append(result, make([]int, c))
		for j := 0; j < c; j++ {
			result[i][j], queue = queue[0], queue[1:]
		}
	}
	return result
}
