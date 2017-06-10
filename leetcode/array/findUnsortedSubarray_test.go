package leetcode

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	if r := findUnsortedSubarray(nums); r != 5 {
		t.Fatal(nums, r)
	}
}
