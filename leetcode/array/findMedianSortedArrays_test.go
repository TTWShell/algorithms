package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	if r := findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}); r != 5.5 {
		t.Fatal(r)
	}

	if r := findMedianSortedArrays([]int{}, []int{2}); r != 2.0 {
		t.Fatal(r)
	}

	if r := findMedianSortedArrays([]int{1, 3}, []int{2}); r != 2.0 {
		t.Fatal(r)
	}

	if r := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); r != 2.5 {
		t.Fatal(r)
	}
}
