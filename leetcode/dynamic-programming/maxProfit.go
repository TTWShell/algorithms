/* https://leetcode.com/problems/best-time-to-buy-and-sell-stock/#/description
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Example 1:
Input: [7, 1, 5, 3, 6, 4]
Output: 5

max. difference = 6-1 = 5 (not 7-1 = 6, as selling price needs to be larger than buying price)
Example 2:
Input: [7, 6, 4, 3, 1]
Output: 0

In this case, no transaction is done, i.e. max profit = 0.

题目意思就是：最多交易一次（买入、卖出为一次交易），找到最大利润。
思路很简单：
    1、找到i之前的最小值；
    2、算下i的局部最优解；
    3、max(局部最优)即为全局最优；
*/

package leetcode

func maxProfit(prices []int) int {
	var min, r int

	if len(prices) > 0 {
		min = prices[0]
	}

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if temp := prices[i] - min; temp > r {
				r = temp
			}
		}
	}
	return r
}
