/* https://leetcode.com/problems/largest-triangle-area/description/
You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

Example:

    Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
    Output: 2
    Explanation:
    The five points are show in the figure below. The red triangle is the largest.

https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/04/1027.png

Notes:

    3 <= points.length <= 50.
    No points will be duplicated.
     -50 <= points[i][j] <= 50.
    Answers within 10^-6 of the true value will be accepted as correct.
*/

package leetcode

import "math"

func largestTriangleArea(points [][]int) float64 {
	area := func(a, b, c []int) float64 {
		pointDistance := func(p1, p2 []int) float64 {
			// https://baike.baidu.com/item/%E4%B8%A4%E7%82%B9%E9%97%B4%E8%B7%9D%E7%A6%BB%E5%85%AC%E5%BC%8F
			return math.Sqrt(math.Pow(float64(p1[0]-p2[0]), 2) + math.Pow(float64(p1[1]-p2[1]), 2))
		}
		side1, side2, side3 := pointDistance(a, b), pointDistance(a, c), pointDistance(b, c)
		s := (side1 + side2 + side3) / 2.0
		return math.Sqrt(s * (s - side1) * (s - side2) * (s - side3))
	}

	max := 0.0
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				if tmp := area(points[i], points[j], points[k]); tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}
