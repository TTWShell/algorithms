package leetcode

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	nums := []int{14, 16}
	results := []bool{false, true}

	for i := range nums {
		if r := isPerfectSquare(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
