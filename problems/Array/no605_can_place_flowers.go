package arr

import (
	"fmt"
)

func CanPlaceFlowers(flowerbed []int, n int) bool {
	canPlace := 0
	i := 0
	if len(flowerbed) == 1 {
		if flowerbed[0] == 1 {
			canPlace = 0
		} else {
			canPlace++
		}
		return canPlace >= n
	}
	for i < len(flowerbed)-1 {
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			fmt.Println(i)
			canPlace += 1
			if canPlace >= n {
				return true
			}
			i += 2
		} else if flowerbed[i] != 0 {
			i += 2
		} else {
			i += 1
		}
	}
	// fmt.Println(i, len(flowerbed))
	if i == len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
		canPlace += 1
	}
	return canPlace >= n
}

func CanPlaceFlowers2(flowerbed []int, n int) bool {
	i := 0
	count := 0
	for i < len(flowerbed) {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			i++
			count++
		}
		if count >= n {
			return true
		}
		i++
	}
	return false
}
