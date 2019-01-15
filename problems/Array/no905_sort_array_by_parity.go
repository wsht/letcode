package arr

func SortArrayByParity(A []int) []int {
	oddIndex := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			A[oddIndex], A[i] = A[i], A[oddIndex]
			oddIndex += 1
		}

	}
	return A
}
