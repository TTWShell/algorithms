/* https://leetcode.com/problems/max-points-on-a-line/
Given n points on a 2D plane, find the maximum number of points
that lie on the same straight line.

Example 1:

Input: [[1,1],[2,2],[3,3]]
Output: 3
Explanation:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
Example 2:

Input: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4
Explanation:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
NOTE: input types have been changed on April 15, 2019.
Please reset to default code definition to get new method signature.
*/

package lmath

func maxPoints(points [][]int) int {
	pointCnt := map[[2]int]int{}
	for _, point := range points {
		pointCnt[[2]int{point[0], point[1]}]++
	}

	if len(pointCnt) <= 2 {
		return len(points) // 独立的两个点，必定全部共线
	}

	uniquePoints := make([][2]int, 0, len(pointCnt))
	for point := range pointCnt {
		uniquePoints = append(uniquePoints, point)
	}

	maxCnt := 2
	for i := 0; i < len(uniquePoints)-2; i++ {
		pos := uniquePoints[i]
		for j := i + 1; j < len(uniquePoints)-1; j++ {
			pointJ := uniquePoints[j]
			curCnt := pointCnt[pos] + pointCnt[pointJ]
			for k := j + 1; k < len(uniquePoints); k++ {
				pointK := uniquePoints[k]
				if (pointJ[0]-pos[0])*(pointK[1]-pos[1]) ==
					(pointK[0]-pos[0])*(pointJ[1]-pos[1]) {
					curCnt += pointCnt[pointK]
				}
			}
			if curCnt > maxCnt {
				maxCnt = curCnt
			}
		}
	}
	return maxCnt
}
