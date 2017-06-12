package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	if r := findMedianSortedArrays([]int{1, 3}, []int{2}); r != 2.0 {
		t.Fatal(r)
	}

	if r := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); r != 2.5 {
		t.Fatal(r)
	}
}
