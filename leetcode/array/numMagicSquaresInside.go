/* https://leetcode.com/problems/magic-squares-in-grid/description/
A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

Given an grid of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).

Example 1:

    Input: [[4,3,8,4],
            [9,5,1,9],
            [2,7,6,2]]
    Output: 1
    Explanation:
    The following subgrid is a 3 x 3 magic square:
    438
    951
    276

    while this one is not:
    384
    519
    762

In total, there is only one magic square inside the given grid.
Note:

    1 <= grid.length <= 10
    1 <= grid[0].length <= 10
    0 <= grid[i][j] <= 15
*/

package leetcode

func numMagicSquaresInside(grid [][]int) int {
	isMagic := func(row, col int) bool {
		// i, j is left-top
		noused := make([]int, 10)
		// 3 row, 3 col,  2 diagonals
		sums := make([]int, 8, 8)
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				tmp := grid[i][j]
				if tmp == 0 || tmp > 9 || noused[tmp] == 1 {
					return false
				}
				noused[tmp] = 1
				sums[j-col] += tmp   // 0-2 idx save sum of row
				sums[i+3-row] += tmp // 3-5 idx save sum of col
			}
		}
		sums[6] = grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]
		sums[7] = grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2]
		for i := 1; i < 8; i++ {
			if sums[i-1] != sums[i] {
				return false
			}
		}
		return true
	}
	row, col, res := len(grid), len(grid[0]), 0
	for i := 0; i <= row-3; i++ {
		for j := 0; j <= col-3; j++ {
			if isMagic(i, j) {
				res++
			}
		}
	}
	return res
}
