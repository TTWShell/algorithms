/* https://leetcode.com/problems/find-the-duplicate-number/description/
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Note:

    You must not modify the array (assume the array is read only).
    You must use only constant, O(1) extra space.
    Your runtime complexity should be less than O(n2).
    There is only one duplicate number in the array, but it could be repeated more than once.

*/

package larray

// https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if fast == slow {
			break
		}
	}

	finder := 0
	for {
		finder = nums[finder]
		slow = nums[slow]
		if finder == slow {
			break
		}
	}

	return slow
}

/*
// Binary Search
func findDuplicate(nums []int) int {
	left, right, count := 1, len(nums)-1, 0 // 1 <= input num <= n, n = len(nums)-1

	for right > left {
		mid := left + (right-left)>>1
		count = mid - left + 1

		for i := 0; i < len(nums); i++ {
			if left <= nums[i] && nums[i] <= mid {
				count--
			}
		}

		if right-left == 1 {
			break
		}

		if count < 0 { // in left -- mid
			right = mid
		} else {
			left = mid
		}
	}

	if count < 0 {
		return left
	}
	return right
}
*/

/*
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
*/
