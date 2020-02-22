/* https://leetcode.com/problems/maximal-rectangle/
Given a 2D binary matrix filled with 0's and 1's,
find the largest rectangle containing only 1's and return its area.

Example:

	Input:
	[
	["1","0","1","0","0"],
	["1","0","1","1","1"],
	["1","1","1","1","1"],
	["1","0","0","1","0"]
	]
	Output: 6
*/

package ldp

func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	dp := make([][][2]int, row, row)
	for i := range dp {
		dp[i] = make([][2]int, col, col) // [2]int = row's length, col's height
	}
	if matrix[0][0] == '1' {
		res = 1
		dp[0][0] = [2]int{1, 1}
	}
	for j := 1; j < col; j++ {
		if matrix[0][j] == '1' {
			tmp := dp[0][j-1][0] + 1
			dp[0][j] = [2]int{tmp, 1}
			res = max(res, tmp)
		}
	}
	for i := 1; i < row; i++ {
		if matrix[i][0] == '1' {
			tmp := dp[i-1][0][1] + 1
			dp[i][0] = [2]int{1, tmp}
			res = max(res, tmp)
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			left, up := matrix[i][j-1], matrix[i-1][j]
			dpLeft, dpUp := dp[i][j-1], dp[i-1][j]
			if left == '0' && up == '0' {
				dp[i][j] = [2]int{1, 1}
			} else if left == '0' {
				dp[i][j] = [2]int{1, dpUp[1] + 1}
			} else if up == '0' {
				dp[i][j] = [2]int{dpLeft[0] + 1, 1}
			} else {
				dp[i][j] = [2]int{dpLeft[0] + 1, dpUp[1] + 1}
			}

			// 只用找up
			tmp := dp[i][j][0]
			res = max(res, max(tmp, dp[i][j][1]))
			for k := i - 1; k >= 0 && dp[k][j][0] != 0; k-- {
				tmp = min(dp[k][j][0], tmp)
				res = max(res, (i-k+1)*tmp)
			}
		}
	}
	return res
}
