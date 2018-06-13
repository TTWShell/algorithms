/* https://leetcode.com/problems/sort-colors/description/
Given an array with n objects colored red, white or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note:
You are not suppose to use the library's sort function for this problem.


Follow up:
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.

Could you come up with an one-pass algorithm using only constant space?
*/

package larray

func sortColors(nums []int) {
	const (
		red  = 0
		blue = 2
	)
	flagRed, flagBlue := 0, len(nums)-1

	for i := 0; i <= flagBlue; i++ {
		for flagRed <= flagBlue && nums[flagRed] == red {
			flagRed++
		}
		for flagBlue >= flagRed && nums[flagBlue] == blue {
			flagBlue--
		}
		if flagRed >= flagBlue {
			break
		}

		if nums[flagRed] == blue || nums[flagBlue] == red {
			nums[flagRed], nums[flagBlue] = nums[flagBlue], nums[flagRed]
			i--
		} else {
			if nums[i] == blue && i < flagBlue {
				nums[i], nums[flagBlue] = nums[flagBlue], nums[i]
			}
			if nums[i] == red && i > flagRed {
				nums[i], nums[flagRed] = nums[flagRed], nums[i]
			}
		}
	}
}
