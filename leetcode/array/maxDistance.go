/* https://leetcode.com/problems/maximum-distance-in-arrays/#/description
Given m arrays, and each array is sorted in ascending order.
Now you can pick up two integers from two different arrays (each array picks one) and calculate the distance.
We define the distance between two integers a and b to be their absolute difference |a-b|. Your task is to find the maximum distance.

Example 1:
    Input:
    [[1,2,3],
     [4,5],
     [1,2,3]]
    Output: 4
    Explanation:
        One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.

Note:
    Each given array will have at least 1 number. There will be at least two non-empty arrays.
    The total number of the integers in all the m arrays will be in the range of [2, 10000].
    The integers in the m arrays will be in the range of [-10000, 10000].
*/

package larray

import "math"

func maxDistance(arrays [][]int) int {
	var (
		minaValue, minaIndex = math.MaxInt32, 0
		minbValue            = math.MaxInt32
		maxaValue            = math.MinInt32
		maxbValue, maxbIndex = math.MinInt32, 0
	)

	for i, array := range arrays {
		curMin, curMax := array[0], array[len(array)-1]
		if curMin < minaValue {
			minbValue = minaValue
			minaValue, minaIndex = curMin, i

		} else if curMin < minbValue {
			minbValue = curMin
		}
		if curMax > maxbValue {
			maxaValue = maxbValue
			maxbValue, maxbIndex = curMax, i
		} else if curMax > maxaValue {
			maxaValue = curMax
		}
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if minaIndex != maxbIndex {
		return abs(maxbValue - minaValue)
	}
	return max(abs(maxbValue-minbValue), abs(maxaValue-minaValue))
}
