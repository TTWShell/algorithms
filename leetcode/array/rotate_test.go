package leetcode

import "testing"

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	result := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(nums, 3)
	for i := 0; i < len(nums); i++ {
		if nums[i] != result[i] {
			t.Error(nums, result)
		}
	}
}
