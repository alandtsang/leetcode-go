// 找到最小值 min 和最小值之后的最大值，profit = max - min
// 如果当前值 - 前值 < min，更新 min
// 如果当前值 - min > profit，更新 profit
package maxprofit1

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min, profit int
	min = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			if prices[i] < min {
				min = prices[i]
			}
		} else {
			if profit < prices[i]-min {
				profit = prices[i] - min
			}
		}
	}
	return profit
}
