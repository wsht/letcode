package arr

func MaxChunksToSorted(arr []int) int {
	chunk := 0
	n := len(arr)
	max := -1
	for i := 0; i < n; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if max == i {
			chunk++
		}
	}
	return chunk
}
