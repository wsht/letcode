package Greedy

import (
	"sort"
)

func FindcontentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	for j := 0; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			i++
		}
	}

	return i
}
