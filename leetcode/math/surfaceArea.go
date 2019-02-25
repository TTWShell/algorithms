/* https://leetcode.com/problems/surface-area-of-3d-shapes/description/
On a N * N grid, we place some 1 * 1 * 1 cubes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Return the total surface area of the resulting shapes.

Example 1:

	Input: [[2]]
	Output: 10

Example 2:

	Input: [[1,2],[3,4]]
	Output: 34

Example 3:

	Input: [[1,0],[0,2]]
	Output: 16

Example 4:

	Input: [[1,1,1],[1,0,1],[1,1,1]]
	Output: 32

Example 5:

	Input: [[2,2,2],[2,1,2],[2,2,2]]
	Output: 46

Note:

	1 <= N <= 50
	0 <= grid[i][j] <= 50
*/

package lmath

func surfaceArea(grid [][]int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	I, J := len(grid), len(grid[0])
	xx, xz, xy := 0, 0, 0
	for i := 0; i < I; i++ {
		tmpXZ, tmpXY := 0, 0
		for j := 0; j < J; j++ {
			if grid[i][j] != 0 {
				xx += 2 // bottom and top
			}
			xz += abs(grid[i][j] - tmpXZ)
			xy += abs(grid[j][i] - tmpXY)
			tmpXZ, tmpXY = grid[i][j], grid[j][i]
		}
		xz += tmpXZ
		xy += tmpXY
	}
	return xx + xz + xy
}
