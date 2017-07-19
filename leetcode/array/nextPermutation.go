/* https://leetcode.com/problems/next-permutation/#/description
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

package leetcode

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	swapIndex := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			swapIndex = i - 1
			break
		}
	}

	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if swapIndex != -1 {
		for i := len(nums) - 1; i > swapIndex; i-- {
			if nums[i] > nums[swapIndex] {
				nums[i], nums[swapIndex] = nums[swapIndex], nums[i]
				reverse(nums[swapIndex+1:])
				break
			}
		}
	} else {
		reverse(nums)
	}
}
