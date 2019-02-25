/* https://leetcode.com/problems/rotting-oranges/description/
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.

Example 1:

https://assets.leetcode.com/uploads/2019/02/16/oranges.png

	Input: [[2,1,1],[1,1,0],[0,1,1]]
	Output: 4

Example 2:

	Input: [[2,1,1],[0,1,1],[1,0,1]]
	Output: -1
	Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:

	Input: [[0,2]]
	Output: 0
	Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.

Note:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] is only 0, 1, or 2.
*/

package larray

func orangesRotting(grid [][]int) int {
	queue := [][2]int{}
	visited := map[[2]int]bool{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
				visited[[2]int{i, j}] = true
			}
		}
	}

	depth := 0
	for len(queue) != 0 {
		tmp := [][2]int{}
		for _, node := range queue {
			for _, cur := range [][2]int{{node[0] - 1, node[1]}, {node[0] + 1, node[1]}, {node[0], node[1] - 1}, {node[0], node[1] + 1}} {
				i, j := cur[0], cur[1]
				if _, ok := visited[[2]int{i, j}]; !ok && 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) && grid[i][j] == 1 {
					tmp = append(tmp, [2]int{i, j})
					visited[[2]int{i, j}] = true
				}
			}
		}
		queue = tmp
		depth++
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if _, ok := visited[[2]int{i, j}]; !ok {
					return -1
				}
			}
		}
	}
	if depth > 0 {
		depth--
	}
	return depth
}
