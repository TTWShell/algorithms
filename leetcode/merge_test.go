package leetcode

import "testing"

func Test_merge(t *testing.T) {
	temp := []int{1, 3, 4, 6}
	nums2 := []int{2, 5, 7, 8, 9}
	nums1 := make([]int, len(temp)+len(nums2))
	copy(nums1, temp)
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	merge(nums1, len(temp), nums2, len(nums2))

	for i := 0; i < len(result); i++ {
		if result[i] != nums1[i] {
			t.Log(result, nums1)
			t.Fail()
		}
	}
}
