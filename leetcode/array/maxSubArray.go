/* https://leetcode.com/problems/maximum-subarray/#/description
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.

More practice:
    If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
思路：直接开始累加，如果累加的和比当前元素还要小，ignore，开始下一轮查找，和保存在i-1的位置。这样，最后nums中保存了所有的局部最优解。
*/
package larray

func maxSubArray(nums []int) int {
	max, len := nums[0], len(nums)

	for i := 1; i < len; i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
