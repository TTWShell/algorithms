package leetcode

import "testing"

func Test_intersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := []int{2}

	r := intersection(nums1, nums2)

	for i := range r {
		if r[i] != result[i] {
			t.Fatal(r, result)
		}
	}
}
