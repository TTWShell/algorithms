/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/#/description
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times). However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

题意：不限制交易次数，计算所有交易的利润和。注意有个前提条件，不能同时参与多个交易，也就是买进前手上要保持无货状态。
思路：
    1、还是要找到i之前的最小值，这个最小值应该是上个交易周期结束后的局部最小值；
    2、只要i-1的利润大于i的利润，我们就应该在i-1的时候卖掉，例如：1，3，2，5；
*/

package ldp

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
