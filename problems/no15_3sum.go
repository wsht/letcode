package main

/**
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

/**
O(n^3)
暴力解决方法
错误 对于重复问题解决
*/
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	MySort(nums)
	for i := 0; i < numsLen-2; i++ {
		for j := i + 1; j < numsLen-1; j++ {
			for k := j + 1; k < numsLen; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					item := []int{nums[i], nums[j], nums[k]}
					result = append(result, item)

					for k+1 < numsLen && nums[k+1] == nums[k] {
						k++
					}
				}
			}
			for j+1 < numsLen && nums[j+1] == nums[j] {
				j++
			}
		}
		for i+1 < numsLen && nums[i+1] == nums[i] {
			i++
		}
	}
	return result
}

func MySort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minNums := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if minNums > nums[j] {
				nums[i] = nums[j]
				nums[j] = minNums
				minNums = nums[i]
			}
		}
	}
}

// func unique(arr []int) []int {
// 	deleteIndex := []int{}
// 	for i := 0; i < len(arr); i++ {
// 		value := arr[i]
// 		fmt.Println(value)
// 		for j := i + 1; j < len(arr); j++ {
// 			if value == arr[j] {
// 				deleteIndex = append(deleteIndex, j)
// 			}
// 		}
// 	}
// 	fmt.Println(deleteIndex)
// 	for _, index := range deleteIndex {
// 		fmt.Println(index)
// 		fmt.Println(arr[:index])
// 		fmt.Println(arr[index+1:])
// 		arr = append(arr[:index], arr[index+1:]...)
// 		fmt.Println(arr)
// 	}
// 	return arr
// }

func ThreeSum2(nums []int) [][]int {
	MySort(nums)
	res := [][]int{}

	// return res
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		front := i + 1
		back := len(nums) - 1

		for front < back {

			sum := nums[front] + nums[back]

			if sum < target {

				front++

			} else if sum > target {

				back--
			} else {

				tmp := []int{nums[i], nums[front], nums[back]}
				res = append(res, tmp)

				//processing duplicates of number 2
				//Rolling the front pointer to the next different number for wards
				for front < back && nums[front] == tmp[1] {

					front++
				}

				//processing duplicates of number 3
				//rolling the back pointer to the next different number backwards
				for front < back && nums[back] == tmp[2] {

					back--
				}
				// front++
				// back--
			}
		}
		//process duplicates of number 1
		for i+1 < len(nums) && nums[i+1] == nums[i] {

			i++
		}
	}

	return res
}

// func main() {
// 	nums := []int{13, 14, 1, 2, -11, -11, -1, 5, -1, -11, -9, -12, 5, -3, -7, -4, -12, -9, 8, -13, -8, 2, -6, 8, 11, 7, 7, -6, 8, -9, 0, 6, 13, -14, -15, 9, 12, -9, -9, -4, -4, -3, -9, -14, 9, -8, -11, 13, -10, 13, -15, -11, 0, -14, 5, -4, 0, -3, -3, -7, -4, 12, 14, -14, 5, 7, 10, -5, 13, -14, -2, -6, -9, 5, -12, 7, 4, -8, 5, 1, -10, -3, 5, 6, -9, -5, 9, 6, 0, 14, -15, 11, 11, 6, 4, -6, -10, -1, 4, -11, -8, -13, -10, -2, -1, -7, -9, 10, -7, 3, -4, -2, 8, -13}
// 	// nums := []int{}
// 	// sort(nums)
// 	// fmt.Println(unique(nums))
// 	fmt.Println(threeSum(nums))
// 	// fmt.Println(threeSum2(nums))
// }
