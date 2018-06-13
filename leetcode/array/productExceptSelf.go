/*https://leetcode.com/problems/product-of-array-except-self/description/
Given an array of n integers where n > 1, nums, return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Solve it without division and in O(n).

For example, given [1,2,3,4], return [24,12,8,6].

Follow up:
Could you solve it with constant space complexity? (Note: The output array does not count as extra space for the purpose of space complexity analysis.)
*/

package larray

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length, length)

	res[0] = nums[0]
	for i := 1; i < length-1; i++ {
		res[i] = nums[i] * res[i-1]
	}

	right := 1
	for i := length - 1; i > 0; i-- {
		res[i] = res[i-1] * right
		right *= nums[i]
	}
	res[0] = right

	return res
}
