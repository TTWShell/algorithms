/* https://leetcode.com/problems/largest-number-at-least-twice-of-others/description/
In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:
    Input: nums = [3, 6, 1, 0]
    Output: 1
    Explanation: 6 is the largest integer, and for every other number in the array x,
    6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
Example 2:
    Input: nums = [1, 2, 3, 4]
    Output: -1
    Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
Note:
    nums will have a length in the range [1, 50].
    Every nums[i] will be an integer in the range [0, 99].
*/

package leetcode

// according to note, can use bucket
func dominantIndex(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}

	bucket := make([]int, 100)
	for i, num := range nums {
		bucket[num] = i + 1
	}

	var largest, largestIndex int
	for i := 99; i >= 0; i-- {
		if bucket[i] == 0 {
			continue
		}
		if largest == 0 {
			largest, largestIndex = i, bucket[i]-1
			continue
		}
		if largest >= 2*i {
			return largestIndex
		}
		break
	}
	return -1
}

/*
func dominantIndex(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}

	var largest, second, largestIndex int
	if nums[0] > nums[1] {
		largest, second, largestIndex = nums[0], nums[1], 0
	} else {
		largest, second, largestIndex = nums[1], nums[0], 1
	}

	for i := 2; i < length; i++ {
		cur := nums[i]
		if cur > largest {
			second, largest, largestIndex = largest, cur, i
		} else if cur > second {
			second = cur
		}
	}

	if largest >= second*2 {
		return largestIndex
	}
	return -1
}
*/
