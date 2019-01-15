package arr

// func MoveZeros(nums []int) {
// 	j := len(nums)
// 	for i := 0; i < j; i++ {
// 		if nums[i] == 0 {
// 			//move to tail
// 			for j > i && nums[j-1] == 0 {
// 				j--
// 			}
// 			if i <= j-1 {
// 				nums[i], nums[j-1] = nums[j-1], nums[i]
// 			}
// 		}
// 	}
// }

func MoveZeros(nums []int) {
	zeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex += 1
		}
	}
}
