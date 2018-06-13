package lht

import "testing"

func Test_intersection(t *testing.T) {
	nums1 := [][]int{{1, 2, 2, 1}, {1, 1}, {}, {1}}
	nums2 := [][]int{{2, 2}, {}, {1, 1}, {1, 1}}
	result := [][]int{{2}, {}, {}, {1}}

	for i := range nums1 {
		r := intersection(nums1[i], nums2[i])

		for j := range r {
			if r[j] != result[i][j] {
				t.Fatal(r, result[i])
			}
		}
	}
}
