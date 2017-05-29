package leetcode

import (
	"sort"
	"testing"
)

func Test_intersect(t *testing.T) {
	nums1 := [][]int{{1, 2, 2, 1}, {1, 1}, {}, {1}, {1, 1}, {4, 7, 9, 7, 6, 7}}
	nums2 := [][]int{{2, 2}, {}, {1, 1}, {1, 1}, {1}, {5, 0, 0, 6, 1, 6, 2, 2, 4}}
	result := [][]int{{2, 2}, {}, {}, {1}, {1}, {4, 6}}

	for i := range nums1 {
		r := intersect(nums1[i], nums2[i])
		sort.Slice(r[:], func(i, j int) bool { return r[i] < r[j] })

		for j := range result[i] {
			if r[j] != result[i][j] {
				t.Fatal(r, result[i])
			}
		}
	}

	// test for sort
	unsorted := []int{2, 1, 3, 5, 4}
	sort.Slice(unsorted[:], func(i, j int) bool { return unsorted[i] < unsorted[j] })
	// t.Log(unsorted)
}
