package arr

func Transpose(A [][]int) [][]int {
	result := [][]int{}
	ALen := len(A)
	for i := 0; i < ALen; i++ {
		for j := 0; j < len(A[i]); j++ {
			if len(result)-1 < j {
				result = append(result, make([]int, ALen))
			}
			result[j][i] = A[i][j]
		}
	}
	return result
}
