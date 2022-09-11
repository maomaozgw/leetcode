// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/

package p188

func maxProfit(k int, prices []int) int {
	var minPrices = make([]int, k+1)
	for i := 0; i <= k; i++ {
		minPrices[i] = 10001
	}
	var maxProfits = make([]int, k+1)
	var l = len(prices)
	for i := 0; i < l; i++ {
		for j := 1; j <= k; j++ {
			minPrices[j] = min(minPrices[j], prices[i]-maxProfits[j-1])
			maxProfits[j] = max(maxProfits[j], prices[i]-minPrices[j])
		}
	}
	return maxProfits[k]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
