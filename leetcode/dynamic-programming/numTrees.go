/* https://leetcode.com/problems/unique-binary-search-trees/description/
Given n, how many structurally unique BST's (binary search trees) that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

package leetcode

/*
n == sum of 1～n as root BST
F(n): numTrees of n
F(i, n): i as root, n node

假设有序列：1,2,3,4,5,6,7
当前i为3，则F(4, 7) = F(2) * F(4)
F(n) = F(1, n) + F(2, n) + ... + F(n, n) = F(0)*F(n-1) + F(1)*F(n-2) + ... + F(n-1)*F(0)

*/
func numTrees(n int) int {
	dp := make([]int, n+1, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
