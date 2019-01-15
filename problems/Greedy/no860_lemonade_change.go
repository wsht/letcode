package Greedy

func LemonadeChange(bills []int) bool {
	store := [3]int{}
	for i := 0; i < len(bills); i++ {
		if !isChangeWithMoney(bills[i], &store) {
			return false
		}
	}
	return true
}

func isChangeWithMoney(money int, store *[3]int) bool {
	switch money {
	case 5:
		store[0]++
		return true
	case 10:
		store[1]++
		store[0]--
		return store[0] >= 0
	case 20:
		store[2]++
		if store[1] >= 1 {
			store[1]--
			store[0]--
			return store[0] >= 0
		}
		//else
		store[0] -= 3
		return store[0] >= 0

	}

	return true
}
