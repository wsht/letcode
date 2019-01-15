package arr

func RemoveDuplicates(nums []int) int {
	rIndex := 1
	n := len(nums)
	if n <= 1 {
		return n
	}
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			if rIndex != i {
				nums[rIndex] = nums[i]
			}
			rIndex++
		}
	}
	return rIndex
}
