package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	targets := []int{5, 2, 7, 0}
	results := []int{2, 1, 4, 0}
	for i := 0; i < len(targets); i++ {
		if r := searchInsert(nums, targets[i]); r != results[i] {
			t.Fatal(nums, targets[i], r)
		}
	}
}
