/* https://leetcode.com/problems/third-maximum-number/#/description
Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:
    Input: [3, 2, 1]
    Output: 1

Explanation: The third maximum is 1.

Example 2:
    Input: [1, 2]
    Output: 2

Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:
    Input: [2, 2, 3, 1]
    Output: 1

Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.

*/

package larray

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func thirdMax(nums []int) int {
	max, second, third := MinInt, MinInt, MinInt
	c := 0
	for _, num := range nums {
		if num == max || num == second || num <= third {
			continue
		}
		if num > max {
			tempAgoMax, tempAgoSecond := max, second
			max, second, third = num, tempAgoMax, tempAgoSecond
		} else if num > second {
			second, third = num, second
		} else if num > third {
			third = num
		}
		c++
	}
	if c > 2 {
		return third
	}
	return max
}
