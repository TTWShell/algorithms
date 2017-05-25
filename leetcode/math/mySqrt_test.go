package leetcode

import "testing"

func Test_mySqrt(t *testing.T) {
	nums := []int{0, 1, 2, 4, 6, 10, 8192}
	results := []int{0, 1, 1, 2, 2, 3, 90}
	for i := 0; i < len(nums); i++ {
		if result := mySqrt(nums[i]); result != results[i] {
			t.Error(nums[i], results[i], result)
		}
	}
}
