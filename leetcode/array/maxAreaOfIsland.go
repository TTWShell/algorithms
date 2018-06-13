/* https://leetcode.com/problems/max-area-of-island/description/
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:
    [[0,0,1,0,0,0,0,1,0,0,0,0,0],
     [0,0,0,0,0,0,0,1,1,1,0,0,0],
     [0,1,1,0,1,0,0,0,0,0,0,0,0],
     [0,1,0,0,1,1,0,0,1,0,1,0,0],
     [0,1,0,0,1,1,0,0,1,1,1,0,0],
     [0,0,0,0,0,0,0,0,0,0,1,0,0],
     [0,0,0,0,0,0,0,1,1,1,0,0,0],
     [0,0,0,0,0,0,0,1,1,0,0,0,0]]
    Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
Example 2:
    [[0,0,0,0,0,0,0,0]]
    Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
*/

package larray

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(grid [][]int, row, col int) (sum int)
	dfs = func(grid [][]int, row, col int) (sum int) {
		sum, grid[row][col] = 1, 2
		if tmp := col - 1; tmp >= 0 && grid[row][tmp] == 1 {
			sum += dfs(grid, row, tmp)
		}
		if tmp := row - 1; tmp >= 0 && grid[tmp][col] == 1 {
			sum += dfs(grid, tmp, col)
		}
		if tmp := col + 1; tmp < len(grid[0]) && grid[row][tmp] == 1 {
			sum += dfs(grid, row, tmp)
		}
		if tmp := row + 1; tmp < len(grid) && grid[tmp][col] == 1 {
			sum += dfs(grid, tmp, col)
		}
		return sum
	}

	res := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				if tmp := dfs(grid, row, col); tmp > res {
					res = tmp
				}
			}
			if grid[row][col] == 2 {
				grid[row][col] = 1
			}
		}
	}
	return res
}
