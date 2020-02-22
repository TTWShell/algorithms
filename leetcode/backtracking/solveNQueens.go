/* https://leetcode.com/problems/n-queens/
The n-queens puzzle is the problem of placing n queens
on an n×n chessboard such that no two queens attack each other.

https://assets.leetcode.com/uploads/2018/10/12/8-queens.png

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of
the n-queens' placement, where 'Q' and '.' both indicate
a queen and an empty space respectively.

Example:

	Input: 4
	Output: [
	[".Q..",  // Solution 1
	"...Q",
	"Q...",
	"..Q."],

	["..Q.",  // Solution 2
	"Q...",
	"...Q",
	".Q.."]
	]
	Explanation: There exist two distinct solutions to the 4-queens
	puzzle as shown above.
*/

package lbacktracking

func solveNQueens(n int) [][]string {
	// 国际象棋中，皇后可以在横、竖、斜线上不限步数地吃掉其他棋子，
	// 所以N个皇后不能在同一行、同一列以及同一斜线上
	res := [][]string{}
	if n < 1 {
		return res
	}

	int2str := func(solution []int) []string {
		tmp, res := make([]byte, n, n), make([]string, n, n)
		for idx, col := range solution {
			for i := 0; i < n; i++ {
				if i != col {
					tmp[i] = '.'
				} else {
					tmp[i] = 'Q'
				}
			}
			res[idx] = string(tmp)
		}
		return res
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

	intRes := dfs([][2]int{}, 0)
	for _, solution := range intRes {
		res = append(res, int2str(solution))
	}
	return res
}
