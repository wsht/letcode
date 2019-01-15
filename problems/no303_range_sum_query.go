package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	return NumArray{sums: sums}
}
func (this *NumArray) SumRange(i int, j int) int {
	return this.sums[j+1] - this.sums[i]
}
