package leetcode

import "testing"

func Test_search(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}

	for i, target := range nums {
		if r := search(nums, target); r != i {
			t.Fatal(nums, target, r)
		}
	}
}
