/* https://leetcode.com/problems/unique-paths/description/
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

https://leetcode.com/static/images/problemset/robot_maze.png

Above is a 3 x 7 grid. How many possible unique paths are there?

Note: m and n will be at most 100.
*/

package leetcode

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	if m == 1 {
		return uniquePaths(1, n-1)
	} else if n == 1 {
		return uniquePaths(m-1, 1)
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}
