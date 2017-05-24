/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/#/description
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times). However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
*/

package leetcode

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, sum, profit := prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			sum += profit
			min, profit = prices[i], 0
		} else {
			if temp := prices[i] - min; temp > profit {
				profit = temp
			}
		}
	}
	return sum + profit
}
