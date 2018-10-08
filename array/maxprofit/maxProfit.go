// 数组分3种情况：
// 1. 递减序列，没有进行交易，利润为 0
// 2. 递增序列，第一天买入，最后一天卖出，利润为 prices[n-1] - prices[0]
// 3. 有增有降，如果当前值比之前值小，min = 当前值，计算 totalSum
// 当前值比之前值大，则需要判断 curSum 和 (当前值 - min) 的大小
package maxprofit

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min, curSum, totalSum int
	min = prices[0] // 第一个元素为 min
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] < 0 { // 后值 < 前值，计算 totalSum
			totalSum += curSum
			curSum = 0
			min = prices[i]
		} else { // 后值 > 前值，判断 curSum 是否需要更新
			if curSum < (prices[i] - min) {
				curSum = prices[i] - min
			}
		}
	}
	return totalSum + curSum
}
