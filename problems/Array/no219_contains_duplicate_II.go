package arr

func ContainsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if i+j > n-1 {
				break
			}
			if nums[i] == nums[i+j] {
				return true
			}
		}
	}
	return false
}

func ContainsDuplicateWithHashmap(nums []int, k int) bool {
	hmap := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		index, ok := hmap[nums[i]]
		if ok && i-index <= k {
			return true
		}
		hmap[nums[i]] = i
	}

	return false
}
