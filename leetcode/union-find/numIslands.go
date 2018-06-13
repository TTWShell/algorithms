/* https://leetcode.com/problems/number-of-islands/description/
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

11110
11010
11000
00000
Answer: 1

Example 2:

11000
11000
00100
00011
Answer: 3
*/

package lunionfind

func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'
		if i > 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if j > 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if i < len(grid)-1 && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j < len(grid[0])-1 && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}

	res := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}

/*
import "github.com/TTWShell/algorithms/data-structure/union-find"

// Union-find solution
func numIslands(grid [][]byte) int {
	uf := unionfind.New()

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union([2]int{i, j}, [2]int{i - 1, j})
				}
				if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
					uf.Union([2]int{i, j}, [2]int{i, j + 1})
				}
				uf.Union([2]int{i, j}, [2]int{i, j})
			}
		}
	}
	return uf.Count()
}
*/
