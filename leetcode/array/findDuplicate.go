/* https://leetcode.com/problems/find-the-duplicate-number/description/
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Note:

    You must not modify the array (assume the array is read only).
    You must use only constant, O(1) extra space.
    Your runtime complexity should be less than O(n2).
    There is only one duplicate number in the array, but it could be repeated more than once.

*/

package leetcode

// O(n^2)
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	panic("input error")
}
