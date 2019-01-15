package arr

func HasGroupsSizseX(deck []int) bool {
	hmap := map[int]int{}
	n := len(deck)
	for i := 0; i < n; i++ {
		hmap[deck[i]] += 1
	}
	// fmt.Println(hmap)
	res := false
	for x := 2; x <= n; x++ {
		if n%x == 0 {
			res = true
			//公约数查找
			for num := range hmap {
				count := hmap[num]
				// fmt.Println(x, count)
				if count%x != 0 {
					res = false
					break
				}
			}
			if res {
				break
			}
		}
	}
	return res
}
