/* https://leetcode.com/problems/minimum-moves-to-equal-array-elements/#/description
Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

Example:

    Input:
    [1,2,3]

    Output:
    3

Explanation:
    Only three moves are needed (remember each move increments two elements):

    [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
*/

package leetcode

func minMoves(nums []int) int {
	// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/#/solutions
	//  sum + m * (n - 1) = x * n  m为加1的次数，x为最终值
	//  x = minNum + m
	// sum - minNum * n = m
	min, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		sum += nums[i]
	}
	return sum - min*len(nums)
}
