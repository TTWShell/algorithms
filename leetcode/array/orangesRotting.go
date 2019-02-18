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
	cur, tmp := make([][]int, len(grid)), make([][]int, len(grid))
	countFresh, countRotten := 0, 0
	for i := 0; i < len(grid); i++ {
		cur[i] = make([]int, len(grid[i]))
		tmp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			cur[i][j] = grid[i][j]
			tmp[i][j] = grid[i][j]
			if grid[i][j] == 1 {
				countFresh++
			} else if grid[i][j] == 2 {
				countRotten++
			}
		}
	}

	if countRotten == 0 && countFresh > 0 {
		return -1
	}
	if countFresh == 0 {
		return 0
	}

	for count := 0; ; count++ {
		countFresh, countRotten = 0, 0
		changed := false
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				switch cur[i][j] {
				case 1:
					countFresh++
					allEmpty := true
					if (i > 0 && tmp[i-1][j] != 0) || (i < len(grid)-1 && tmp[i+1][j] != 0) || (j > 0 && tmp[i][j-1] != 0) || (j < len(grid[i])-1 && tmp[i][j+1] != 0) {
						allEmpty = false
					}
					if allEmpty {
						return -1
					}
				case 2:
					countRotten++
					if i > 0 && tmp[i-1][j] == 1 {
						changed = true
						tmp[i-1][j] = 2
					}
					if i < len(grid)-1 && tmp[i+1][j] == 1 {
						changed = true
						tmp[i+1][j] = 2
					}
					if j > 0 && tmp[i][j-1] == 1 {
						changed = true
						tmp[i][j-1] = 2
					}
					if j < len(grid[i])-1 && tmp[i][j+1] == 1 {
						changed = true
						tmp[i][j+1] = 2
					}
				}
			}
		}
		if changed == false {
			if countFresh > 0 {
				return -1
			}
			count--
			return count
		}
		cur, tmp = tmp, cur
	}
}
