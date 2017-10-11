/* https://leetcode.com/problems/triangle/description/
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle
    [
         [2],
        [3,4],
       [6,5,7],
      [4,1,8,3]
    ]
    The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
*/

package leetcode

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if len(triangle) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, length)
	copy(dp, triangle[length-1])

	for i := length - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
