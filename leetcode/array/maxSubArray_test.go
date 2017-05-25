package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if r := maxSubArray(nums); r != 6 {
		t.Error(r)
	}

	nums = []int{-2, -3, -1, -5}
	if r := maxSubArray(nums); r != -1 {
		t.Error(r)
	}

	nums = []int{-1, 2}
	if r := maxSubArray(nums); r != 2 {
		t.Error(r)
	}
}
