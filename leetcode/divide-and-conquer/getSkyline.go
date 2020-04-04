/* https://leetcode.com/problems/the-skyline-problem/
A city's skyline is the outer contour of the silhouette formed by all the
 buildings in that city when viewed from a distance.
 Now suppose you are given the locations and height of all the buildings
 as shown on a cityscape photo (Figure A), write a program to output
 the skyline formed by these buildings collectively (Figure B).

https://assets.leetcode.com/uploads/2018/10/22/skyline1.png

The geometric information of each building is represented by a triplet of
integers [Li, Ri, Hi], where Li and Ri are the x coordinates of the left
 and right edge of the ith building, respectively, and Hi is its height.
 It is guaranteed that 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX,
 and Ri - Li > 0. You may assume all buildings are perfect rectangles
 grounded on an absolutely flat surface at height 0.

For instance, the dimensions of all buildings in Figure A are recorded as:
[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] .

The output is a list of "key points" (red dots in Figure B) in the format
of [ [x1,y1], [x2, y2], [x3, y3], ... ] that uniquely defines a skyline.
 A key point is the left endpoint of a horizontal line segment.
 Note that the last key point, where the rightmost building ends,
 is merely used to mark the termination of the skyline,
 and always has zero height. Also, the ground in between any two adjacent
  buildings should be considered part of the skyline contour.

For instance, the skyline in Figure B should be represented as:
[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ].

Notes:

	The number of buildings in any input list is guaranteed to be in the
	     range [0, 10000].
	The input list is already sorted in ascending order by the left x
	     position Li.
    The output list must be sorted by the x position.
	There must be no consecutive horizontal lines of equal height in the
	    output skyline. For instance, [...[2 3], [4 5], [7 5], [11 5],
	    [12 7]...] is not acceptable; the three lines of height 5 should be
	    merged into one in the final output as such: [...[2 3], [4 5], [12 7], ...]
*/

package ldc

func getSkyline(buildings [][]int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var merge func(buildings [][]int, start int, end int) [][]int
	merge = func(buildings [][]int, start int, end int) [][]int {
		//return when array is empty
		if start > end {
			return [][]int{}
		}

		if start == end {
			return [][]int{
				[]int{buildings[start][0], buildings[start][2]},
				[]int{buildings[start][1], 0},
			}
		}

		mid := start + (end-start)/2
		left, right := merge(buildings, start, mid), merge(buildings, mid+1, end)

		//left height and right height for the case when two skylines overlap
		leftH, rightH, res := 0, 0, [][]int{}
		for len(left) > 0 && len(right) > 0 {
			x, x1, x2 := 0, left[0][0], right[0][0]
			if x1 < x2 { // when left edge is before the right edge
				x, leftH = left[0][0], left[0][1]
				left = left[1:]
			} else if x1 > x2 { // when left edge is ahead of right edge
				x, rightH = right[0][0], right[0][1]
				right = right[1:]
			} else { // when two vertical edges overlap
				x, leftH, rightH = left[0][0], left[0][1], right[0][1]
				// both array needs to pop, as we don't want the same kep point appears twice
				left, right = left[1:], right[1:]
			}
			h := max(leftH, rightH)

			//only add new key point when the new key point has a different height, this handles the edge case where two squares starts on the same edge, only the higher one will be added once.
			if len(res) == 0 || h != res[len(res)-1][1] {
				res = append(res, []int{x, h})
			}
		}

		// keep appending the left or right when the other array runs out.
		for i := 0; i < len(left); i++ {
			if len(res) == 0 || left[i][1] != res[len(res)-1][1] {
				res = append(res, []int{left[i][0], left[i][1]})
			}
		}
		for j := 0; j < len(right); j++ {
			if len(res) == 0 || right[j][1] != res[len(res)-1][1] {
				res = append(res, []int{right[j][0], right[j][1]})
			}
		}
		return res
	}

	return merge(buildings, 0, len(buildings)-1)
}

// import (
// 	"sort"
// )
//
// 100ms
// func getSkyline(buildings [][]int) [][]int {
// 	if len(buildings) == 0 {
// 		return [][]int{}
// 	}

// 	tmp := map[int]int{}
// 	for _, build := range buildings {
// 		if _, ok := tmp[build[0]]; !ok {
// 			tmp[build[0]] = build[2]
// 		}
// 		if _, ok := tmp[build[1]]; !ok {
// 			tmp[build[1]] = 0
// 		}
// 	}

// 	points := make([][]int, 0, len(tmp))
// 	for x, h := range tmp {
// 		points = append(points, []int{x, h})
// 	}
// 	sort.Slice(points, func(i, j int) bool {
// 		return points[i][0] < points[j][0]
// 	})

// 	left := 0
// 	for _, build := range buildings {
// 		for points[left][0] < build[0] {
// 			left++
// 		}
// 		for idx := left; points[idx][0] < build[1]; idx++ {
// 			if points[idx][1] < build[2] {
// 				points[idx][1] = build[2]
// 			}
// 		}
// 	}

// 	res := [][]int{points[0]}
// 	for i := 1; i < len(points); i++ {
// 		if points[i][1] != res[len(res)-1][1] {
// 			res = append(res, points[i])
// 		}
// 	}
// 	return res
// }
