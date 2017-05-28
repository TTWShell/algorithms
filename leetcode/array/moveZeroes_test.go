package leetcode

import "testing"

func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	result := []int{1, 3, 12, 0, 0}

	moveZeroes(nums)
	for i := range nums {
		if nums[i] != result[i] {
			t.Fatal(nums, result)
		}
	}
}
