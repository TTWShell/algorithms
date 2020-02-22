/* https://leetcode.com/problems/jump-game-ii/
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:

	Input: [2,3,1,1,4]
	Output: 2
	Explanation: The minimum number of jumps to reach the last index is 2.
		Jump 1 step from index 0 to 1, then 3 steps to the last index.

Note:

	You can assume that you can always reach the last index.
*/

package larray

func jump(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}

	steps := 0
	for i := 0; i < length; {
		if nums[i]+i+1 >= length {
			return steps + 1
		}

		steps++
		next := i + 1
		for j := i + 1; j <= i+nums[i]; j++ {
			if nums[next]-nums[j] < j-next { // 找到下一跳可以跳最远的, greedy
				next = j
			}
		}
		i = next
	}
	return 0
}

// func jump(nums []int) int {
// 	length := len(nums)
// 	steps := make([]int, length, length)

// 	for i, num := range nums {
// 		for idx := 1; idx <= num; idx++ {
// 			next := i + idx
// 			if next >= length {
// 				break
// 			}
// 			if steps[next] == 0 || steps[next] > steps[i]+1 {
// 				steps[next] = steps[i] + 1
// 			}
// 		}
// 		if steps[length-1] != 0 {
// 			break
// 		}
// 	}
// 	return steps[length-1]
// }
