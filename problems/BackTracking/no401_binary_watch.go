package backtrack

import (
	"math"
	"math/bits"
	"strconv"
)

func ReadBinaryWatch(nums int) []string {
	ret := []string{}
	hmap := map[string]int{}
	binarywatchHelper(&ret, hmap, nums, 0, 0, 0, 0)
	return ret
}

func binarywatchHelper(ret *[]string, hmap map[string]int, nums, ms, hs, hour, minute int) {

	if hour >= 12 || minute >= 60 {
		return
	}
	if nums == 0 {

		// fmt.Println(hour)
		h := strconv.Itoa(hour)
		// fmt.Println(h)
		m := ""
		if minute < 10 {
			m = "0" + strconv.Itoa(minute)
		} else {
			m = strconv.Itoa(minute)
		}
		time := h + ":" + m
		_, ok := hmap[time]
		if ok {
			return
		}
		hmap[time] = 1
		*ret = append(*ret, h+":"+m)
		return
	}
	for h := hs; h < 4; h++ {
		binarywatchHelper(ret, hmap, nums-1, ms, h+1, hour+int(math.Pow(2.0, float64(h))), minute)
	}

	for m := ms; m < 6; m++ {
		binarywatchHelper(ret, hmap, nums-1, m+1, hs, hour, minute+int(math.Pow(2.0, float64(m))))
	}
}

func ReadBinaryWatchByBinManipulation(num int) []string {
	ret := []string{}
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			bit := h<<6 | m
			if bits.OnesCount(uint(bit)) == num {
				hour := strconv.Itoa(h)
				minute := strconv.Itoa(m)
				if m < 10 {
					minute = "0" + strconv.Itoa(m)
				}
				ret = append(ret, hour+":"+minute)
			}
		}
	}
	return ret
}
