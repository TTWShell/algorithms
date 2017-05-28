package leetcode

import "testing"

func Test_canWinNim(t *testing.T) {
	nums := []int{4}
	results := []bool{false}

	for i := range nums {
		if r := canWinNim(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
