package backtrack

func Partition(s string) [][]string {
	ret := [][]string{}
	list := []string{}
	partitionHelper(&ret, s, list, len(s))
	return ret
}

func partitionHelper(ret *[][]string, s string, list []string, pos int) {
	// fmt.Println(ret, s, list, pos)
	if pos == 0 {
		listcp := make([]string, len(list))
		copy(listcp, list)
		*ret = append(*ret, listcp)
		return
	}

	for i := 1; i <= len(s); i++ {
		if isPartition(s[0:i]) {
			list = append(list, s[0:i])
			// fmt.Println(list, pos, i)
			partitionHelper(ret, s[i:], list, pos-i)
			list = list[0 : len(list)-1]
		}
	}
}

func isPartition(a string) bool {
	end := len(a) - 1
	start := 0
	// fmt.Println(a)
	for start <= end {
		if a[end] != a[start] {
			return false
		}
		start++
		end--
	}

	return true
}
