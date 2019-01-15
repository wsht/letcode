package arr

func RemoveElement(nums []int, val int) int {
	resLen := 0
	i := 0
	j := len(nums)
	for i <= j-1 {
		if nums[i] != val {
			// fmt.Println("!=", nums[i], val)
			resLen++
		} else {
			for j >= 1 && nums[j-1] == val {
				j--
			}
			if i < j-1 {
				nums[i], nums[j-1] = nums[j-1], nums[i]
				// fmt.Println("exchange", i, nums[i], val)
				resLen++
				j--
			}
		}
		i++
	}

	return resLen
}

func RemoveElement2(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}
