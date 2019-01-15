package arr

//the input arr is sort

func TwoSum(numbers []int, target int) []int {
	low, hight := 0, len(numbers)-1

	for low < hight {
		if numbers[low]+numbers[hight] == target {
			break
		}
		if numbers[low]+numbers[hight] > target {
			hight--
		}
		if numbers[low]+numbers[hight] < target {
			low++
		}
	}

	return []int{low + 1, hight + 1}
}
