/* https://leetcode.com/problems/maximal-square/description/
Given a 2D binary matrix filled with 0's and 1's, find the largest square containing only 1's and return its area.

For example, given the following matrix:

    1 0 1 0 0
    1 0 1 1 1
    1 1 1 1 1
    1 0 0 1 0
Return 4.
*/

package ldp

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	row, col := len(matrix), len(matrix[0])

	dp, side := make([][]int, row), 0
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for j := 0; j < col; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			side = 1
		}
	}
	for i := 1; i < row; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			side = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if dp[i-1][j] == 0 || dp[i][j-1] == 0 || dp[i-1][j-1] == 0 {
				dp[i][j] = 1
			} else if tmp := dp[i-1][j]; tmp == dp[i][j-1] {
				if dp[i-tmp][j-tmp] != 0 {
					dp[i][j] = tmp + 1
				} else {
					dp[i][j] = tmp
				}
			} else {
				min := tmp
				if dp[i][j-1] < min {
					min = dp[i][j-1]
				}
				dp[i][j] = min + 1
			}

			if side < dp[i][j] {
				side = dp[i][j]
			}
		}
	}
	return side * side
}
