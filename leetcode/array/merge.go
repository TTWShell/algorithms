/* https://leetcode.com/problems/merge-sorted-array/#/description
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
    You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
关键的点在于题目已申明，nums1的空间是足够的，所以应该是nums1本身空间就是m+n。
*/

package larray

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	for m, n = m-1, n-1; m >= 0 && n >= 0; k-- {
		if nums1[m] > nums2[n] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
	}
	for n >= 0 {
		nums1[k] = nums2[n]
		k--
		n--
	}
}
