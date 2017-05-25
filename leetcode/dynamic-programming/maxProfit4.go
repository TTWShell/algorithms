/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/#/description
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

具体解法见系列3。
*/

package leetcode

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	if k >= n/2 {
		n, profit := len(prices), 0
		for i := 1; i < n; i++ {
			// as long as there is a price gap, we gain a profit.
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	g := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		g[i] = make([]int, n)
	}
	for i := 1; i <= k; i++ {
		tmpMax := -prices[0]
		for j := 1; j < n; j++ {
			g[i][j] = max(g[i][j-1], prices[j]+tmpMax)
			tmpMax = max(tmpMax, g[i-1][j-1]-prices[j])
		}
	}
	return g[k][n-1]
}
