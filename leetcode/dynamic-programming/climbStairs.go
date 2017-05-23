/* https://leetcode.com/problems/climbing-stairs/#/description
ou are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

假设梯子有n层，那么如何爬到第n层呢，因为每次只能爬1或2步，那么爬到第n层的方法要么是从第n-1层一步上来的，要不就是从n-2层2步上来的，所以递推公式非常容易的就得出了：
dp[n] = dp[n-1] + dp[n-2] （上第n层只有两种解法）
如果梯子有1层或者2层，dp[1] = 1, dp[2] = 2，如果梯子有0层，自然dp[0] = 0
*/
package leetcode

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return n
	}

	step1, step2, r := 2, 1, 0
	for i := 2; i < n; i++ {
		r = step1 + step2
		step2, step1 = step1, r
	}
	return r
}
