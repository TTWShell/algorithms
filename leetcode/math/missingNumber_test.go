package lmath

import "testing"

func Test_missingNumber(t *testing.T) {
	nums := [][]int{{0, 1, 3}, {0, 1, 2, 3, 4, 5, 6, 7, 9}}
	results := []int{2, 8}

	for i := range nums {
		if r := missingNumber(nums[i]); r != results[i] {
			t.Fatal(nums[i], r)
		}
	}
}
