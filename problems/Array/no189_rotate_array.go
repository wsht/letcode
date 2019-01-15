package arr

import (
	"fmt"
)

func Rotate(nums []int, k int) {
	n := len(nums)
	step := k
	if n == 0 {
		return
	}
	step = k % n
	if step == 0 {
		return
	}
	fmt.Println(step, k, n)
	//
	// 	for i := 0; i < n/2; i++ {

	// 	}
	// } else {
	changeIndex := step
	tmp := nums[0]
	i := 0
	for i = 0; i < n; i++ {
		tmpTmp := nums[changeIndex]
		nums[changeIndex] = tmp
		fmt.Println(tmp, changeIndex, nums)
		if i != 0 && changeIndex == step {
			break
		}
		tmp = tmpTmp
		changeIndex = (changeIndex + step) % n
		// fmt.Println(nums)
	}
	fmt.Println(i)
	if i < n/2 {
		changeIndex = step + 1
		tmp = nums[1]
		for i = 0; i < n; i++ {
			tmpTmp := nums[changeIndex]
			nums[changeIndex] = tmp
			fmt.Println(tmp, changeIndex, nums)
			if i != 0 && changeIndex == step+1 {
				break
			}
			tmp = tmpTmp
			changeIndex = (changeIndex + step) % n
			// fmt.Println(nums)
		}
	}

	fmt.Println(tmp, changeIndex, nums)
	// }

	// return

}

func Rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0
	for i := 0; count < n; i++ {
		current := i
		prev := nums[i]
		for true {
			next := (current + k) % n
			tmp := nums[next]
			nums[next] = prev
			// fmt.Println(next, tmp, nums)
			prev = tmp
			current = next
			count++
			if current == i {
				break
			}
		}
	}
	// fmt.Println(nums)
}
