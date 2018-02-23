/* https://leetcode.com/problems/rectangle-area/description/
Find the total area covered by two rectilinear rectangles in a 2D plane.

Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.

https://leetcode.com/static/images/problemset/rectangle_area.png

Assume that the total area is never beyond the maximum possible value of int.
*/

package leetcode

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	edge := func(A, C int) int {
		// 相交或者不相交的情况
		return max(0, C-A)
	}

	area := func(A, B, C, D int) int {
		return edge(A, C) * edge(B, D)
	}

	return area(A, B, C, D) + area(E, F, G, H) - area(max(A, E), max(B, F), min(C, G), min(D, H))
}
