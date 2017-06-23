/* https://leetcode.com/problems/container-with-most-water/#/description
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

package leetcode

func maxArea(height []int) int {
	// x轴上在1,2,...,n点上有许多垂直的线段，长度依次是a1, a2, ..., an
	// 找出两条线段，使他们和x抽围成的面积最大。面积公式是 Min(ai, aj) X |j - i|
	area := func(hi, hj, i, j int) int {
		if hi < hj {
			return hi * (j - i)
		}
		return hj * (j - i)
	}

	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		curArea := area(height[i], height[j], i, j)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if curArea > maxArea {
			maxArea = curArea
		}
	}
	return maxArea
}
