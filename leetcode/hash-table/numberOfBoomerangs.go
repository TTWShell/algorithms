/* https://leetcode.com/problems/number-of-boomerangs/#/description
Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k)
such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs.
You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:
    Input:
    [[0,0],[1,0],[2,0]]

    Output:
    2

    Explanation:
        The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
*/

package leetcode

func numberOfBoomerangs(points [][]int) int {
	count, n := 0, len(points)
	for i := 0; i < n; i++ {
		maps := make(map[int]int, n-i)
		for j := 0; j < n; j++ {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			dis := x*x + y*y
			count += maps[dis] * 2
			maps[dis]++
		}
	}
	return count
}
