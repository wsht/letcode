package tn

import (
	"fmt"
	"math"
	"sort"
)

type SegmentTree struct {
	N    int
	H    int
	tree []int
	lazy []int
}

func NewSegmentTree(N int) SegmentTree {
	H := 1
	//计算树的高度 = log2^N
	for (1 << uint(H)) < N {
		H++
	}
	tree := make([]int, 2*N)
	lazy := make([]int, N)

	return SegmentTree{N, H, tree, lazy}
}

func (segT *SegmentTree) apply(x, val int) {
	segT.tree[x] = int(math.Max(float64(segT.tree[x]), float64(val)))
	if x < segT.N {
		segT.lazy[x] = int(math.Max(float64(segT.lazy[x]), float64(val)))
	}
	fmt.Println(segT.tree, segT.lazy)
}

func (segT *SegmentTree) pull(x int) {
	for x > 1 {
		x >>= 1
		segT.tree[x] = int(math.Max(float64(segT.tree[x*2]), float64(segT.tree[x*2+1])))
		segT.tree[x] = int(math.Max(float64(segT.tree[x]), float64(segT.lazy[x])))
	}
}

func (segT *SegmentTree) push(x int) {
	for h := segT.H; h > 0; h-- {
		y := x >> uint(h)
		if segT.lazy[y] > 0 {
			segT.apply(y*2, segT.lazy[y])
			segT.apply(y*2+1, segT.lazy[y])
			segT.lazy[y] = 0
		}

	}
}

func (segT *SegmentTree) update(L, R, h int) {
	L += segT.N
	R += segT.N
	L0, R0 := L, R
	for L <= R {
		//奇数
		if (L & 1) == 1 {
			segT.apply(L, h)
			L++
		}
		//偶数
		if (R & 1) == 0 {
			segT.apply(R, h)
			R--
		}
		L >>= 1
		R >>= 1
	}
	segT.pull(L0)
	segT.pull(R0)
}

func (segT *SegmentTree) query(L, R int) int {
	L += segT.N
	R += segT.N
	ans := 0
	segT.push(L)
	segT.push(R)
	for L <= R {
		if (L & 1) == 1 {
			ans = int(math.Max(float64(ans), float64(segT.tree[L])))
			L++
		}

		if (R & 1) == 0 {
			ans = int(math.Max(float64(ans), float64(segT.tree[R])))
			R--
		}
		L >>= 1
		R >>= 1
	}
	return ans
}

func coordinateCompression(positions [][]int) (map[int]int, int) {
	list := []int{}
	lmap := map[int]int{}
	for i := 0; i < len(positions); i++ {
		_, ok := lmap[positions[i][0]]
		if !ok {
			list = append(list, positions[i][0])
			lmap[positions[i][0]] = 1
		}
		endpoint := positions[i][0] + positions[i][1] - 1
		_, ok = lmap[endpoint]
		if !ok {
			list = append(list, endpoint)
			lmap[endpoint] = 1
		}
	}

	fmt.Println(lmap)

	sort.Ints(list)
	ret := map[int]int{}
	for i := 0; i < len(list); i++ {
		ret[list[i]] = i
	}
	return ret, len(list)
}

//[[1,2], [2,3], [6,1]]
//1,3,2,5,6
func FallingSquares(positions [][]int) []int {
	index, size := coordinateCompression(positions)
	// fmt.Println(index, size)
	tree := NewSegmentTree(size)
	best := 0
	ans := []int{}
	for i := 0; i < len(positions); i++ {
		L := index[positions[i][0]]
		R := index[positions[i][0]+positions[i][1]-1]

		h := tree.query(L, R) + positions[i][1]
		tree.update(L, R, h)
		best = int(math.Max(float64(best), float64(h)))
		ans = append(ans, best)
	}
	return ans
}

// 萤火虫
// http://www.and2girls.com/index/qrcode?active=badge&channel_code=yhc2&badge_id=31
// ido
// http://www.and2girls.com/index/qrcode?active=badge&channel_code=ido2&badge_id=35
// 新年
// http://www.and2girls.com/index/qrcode?active=badge&channel_code=ny&badge_id=39
