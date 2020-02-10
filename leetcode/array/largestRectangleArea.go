/* https://leetcode.com/problems/largest-rectangle-in-histogram/
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
find the area of largest rectangle in the histogram.

https://assets.leetcode.com/uploads/2018/10/12/histogram.png

Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].

The largest rectangle is shown in the shaded area, which has area = 10 unit.

https://assets.leetcode.com/uploads/2018/10/12/histogram_area.png

Example:

	Input: [2,1,5,6,2,3]
	Output: 10
*/

package larray

func largestRectangleArea(heights []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	length := len(heights)
	res := 0
	for i := 0; i < length; i++ {
		if i+1 < length && heights[i] <= heights[i+1] {
			continue
		}
		minH := heights[i]
		for j := i; j >= 0; j-- {
			minH = min(minH, heights[j])
			area := minH * (i - j + 1)
			res = max(res, area)
		}
	}
	return res
}
