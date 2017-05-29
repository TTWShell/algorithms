package leetcode

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	nums := []int{1, 4, 5, 16, 17}
	results := []bool{true, true, false, true, false}

	for i := range nums {
		if r := isPowerOfFour(nums[i]); r != results[i] {
			t.Fatal(nums[i], results[i])
		}
	}
}
