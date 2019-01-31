/* https://leetcode.com/problems/maximize-distance-to-closest-person/
In a row of seats, 1 represents a person sitting in that seat, and 0 represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

Example 1:

	Input: [1,0,0,0,1,0,1]
	Output: 2
	Explanation:
	If Alex sits in the second open seat (seats[2]), then the closest person has distance 2.
	If Alex sits in any other open seat, the closest person has distance 1.
	Thus, the maximum distance to the closest person is 2.

Example 2:

	Input: [1,0,0,0]
	Output: 3
	Explanation:
	If Alex sits in the last seat, the closest person is 3 seats away.
	This is the maximum distance possible, so the answer is 3.

Note:

	1 <= seats.length <= 20000
	seats contains only 0s or 1s, at least one 0, and at least one 1.
*/

package larray

func maxDistToClosest(seats []int) int {
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	res, left, right := 0, -1, 0
	for i := 0; i < len(seats); i++ {
		right = i
		if seats[i] == 1 {
			if left == -1 {
				res = right
			} else {
				res = max(res, (right-left)/2)
			}
			left = i
		}
	}
	return max(res, right-left)
}
