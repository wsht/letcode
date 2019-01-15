package arr

import (
	"math"
)

func MaxDistToClosest(seats []int) int {
	max := 0
	count := 0
	// start := -1
	tmpStart := -1

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			count++
			tmp := 0
			if tmpStart == -1 || i == len(seats)-1 {
				tmp = count
			} else {
				tmp = count/2 + count%2
			}
			if tmp > max {
				// start = tmpStart
				max = tmp
			}
		} else {
			tmpStart = i
			count = 0
		}
	}
	// fmt.Println(tmpStart, max)
	return max
}

func MaxDistToClosest2(seats []int) int {
	i, j, res := 0, 0, 0
	n := len(seats)
	for ; j < n; j++ {
		if seats[j] == 0 {
			res = int(math.Max(float64(res), float64(j-i)))
		} else {
			res = int(math.Max(float64(res), float64((j-i+1)/2)))
		}
		i = j + 1
	}
	res = int(math.Max(float64(res), float64(n-i)))
	return res
}
