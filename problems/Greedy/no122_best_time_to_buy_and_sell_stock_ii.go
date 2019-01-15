package Greedy

func MaxPrfitII(prices []int) int {
	return calcuate(prices, 0)
}

func calcuate(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}

	max := 0
	for i := 0; i < len(prices); i++ {
		maxProfit := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				profit := calcuate(prices, j+1) + prices[j] - prices[i]
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}
		if maxProfit > max {
			max = maxProfit
		}
	}
	return max
}

func MaxPrfitII_2(prices []int) int {
	n := len(prices)
	i := 0
	valley := prices[0]
	peak := prices[0]
	maxprofit := 0
	for i < n-1 {
		//找最低点
		for i < n-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		//找最高点
		for i < n-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxprofit += peak - valley
	}

	return maxprofit
}

//通过2我们可以知道，total profit 可以是每一次增幅的总和
func MaxPrfitII_3(prices []int) int {
	maxprofit := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}

	return maxprofit
}

//go slice 倒叙访问比正序访问快？
func MaxPrfitII4(prices []int) int {
	maxprofit := 0
	n := len(prices)
	for i := n - 1; i > 0; i-- {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}
