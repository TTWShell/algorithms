package leetcode

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	nums := []int{3, 6, 9, 27, 102}
	results := []bool{true, false, true, true, false}

	for i := range nums {
		if r := isPowerOfThree(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
