package lmath

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	nums := []int{1, 2, 4, 15, 16}
	results := []bool{true, true, true, false, true}

	for i := range nums {
		if r := isPowerOfTwo(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
