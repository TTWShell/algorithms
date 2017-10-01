/* https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.
*/

package leetcode

func removeDuplicates2(nums []int) int {
	length, end := len(nums), 1
	if length <= 2 {
		return length
	}

	for i := 2; i < length; i++ {
		if nums[i] != nums[end-1] {
			end++
			nums[end] = nums[i]
		}
	}
	return end + 1
}
