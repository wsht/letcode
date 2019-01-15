package Greedy

func MinDeletionSize(A []string) int {
	r := len(A)
	c := len(A[0])
	count := 0
	for i := 0; i < c; i++ {
		for j := 1; j < r; j++ {
			if A[j][i] < A[j-1][i] {
				count++
				break
			}
		}
	}

	return count
}
