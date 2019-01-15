package arr

func GetRow(rowIndex int) []int {
	A := make([]int, rowIndex+1)
	A[0] = 1
	for i := 1; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			A[j] += A[j-1]
		}
		// fmt.Println(i, A)
	}
	return A
}
