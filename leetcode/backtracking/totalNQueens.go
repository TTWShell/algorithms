/* https://leetcode.com/problems/n-queens-ii/
The n-queens puzzle is the problem of placing n queens on an
n×n chessboard such that no two queens attack each other.

https://assets.leetcode.com/uploads/2018/10/12/8-queens.png

Given an integer n, return the number of distinct solutions
to the n-queens puzzle.

Example:

	Input: 4
	Output: 2
	Explanation: There are two distinct solutions to the 4-queens
	puzzle as shown below.
	[
	[".Q..",  // Solution 1
	"...Q",
	"Q...",
	"..Q."],

	["..Q.",  // Solution 2
	"Q...",
	"...Q",
	".Q.."]
	]
*/

package lbacktracking

func totalNQueens(n int) int {
	// 同 https://leetcode.com/problems/n-queens/
	if n < 1 {
		return 0
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	search := func(points [][2]int, row int) []int {
		// give used points and target row, return all not used cols
		res, tmp := []int{}, make([]int, n, n)

		for _, point := range points {
			// 列已经被用的，直接标记不可用
			tmp[point[1]] = 1
		}
		for col := 0; col < n; col++ {
			if tmp[col] == 1 {
				continue
			}
			for _, point := range points {
				if abs(point[0]-row) == abs(point[1]-col) {
					tmp[col] = 1
				}
			}
		}

		for col, used := range tmp {
			if used == 0 {
				res = append(res, col)
			}
		}
		return res
	}

	var dfs func(points [][2]int, row int) [][]int
	dfs = func(points [][2]int, row int) [][]int {
		res := [][]int{}
		cols := search(points, row)
		if len(cols) == 0 { // no solutions
			return res
		}
		if row == n-1 { // last row
			return [][]int{cols}
		}

		for _, col := range cols {
			points = append(points, [2]int{row, col})
			for _, tmp := range dfs(points, row+1) {
				res = append(res, append([]int{col}, tmp...))
			}
			points = points[:len(points)-1]
		}
		return res
	}

	return len(dfs([][2]int{}, 0))
}
