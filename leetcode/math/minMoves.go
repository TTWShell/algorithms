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

package lmath

func minMoves(nums []int) int {
	// 给n-1个数字加1，效果等同于给那个未被选中的数字减1 ----> 所有数字都减小到最小值
	// Find the minimum
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	// Sum the diff with min
	var moves int
	for _, num := range nums {
		moves += num - min
	}
	return moves
}

/*
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
*/
