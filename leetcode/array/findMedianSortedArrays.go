/* https://leetcode.com/problems/median-of-two-sorted-arrays/#/description

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
    nums1 = [1, 3]
    nums2 = [2]

    The median is 2.0

Example 2:
    nums1 = [1, 2]
    nums2 = [3, 4]

    The median is (2 + 3)/2 = 2.5
*/

package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// https://discuss.leetcode.com/topic/16797/very-concise-o-log-min-m-n-iterative-solution-with-detailed-explanation
	n1, n2 := len(nums1), len(nums2)
	if n1 < n2 {
		// Make sure nums2 is the shorter one.
		return findMedianSortedArrays(nums2, nums1)
	}

	if n2 == 0 {
		return float64(nums1[(n1-1)/2]+nums1[n1/2]) / 2.0
	}

	lo, hi := 0, n2*2
	for lo <= hi {
		mid2 := (lo + hi) / 2  // Try Cut 2
		mid1 := n1 + n2 - mid2 // Calculate Cut 1 accordingly

		// Get L1, R1, L2, R2 respectively
		L1, L2, R1, R2 := math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
		if mid1 != 0 {
			L1 = nums1[(mid1-1)/2]
		}
		if mid2 != 0 {
			L2 = nums2[(mid2-1)/2]
		}
		if mid1 != n1*2 {
			R1 = nums1[(mid1)/2]
		}
		if mid2 != n2*2 {
			R2 = nums2[(mid2)/2]
		}

		if L1 > R2 {
			lo = mid2 + 1 // A1's lower half is too big; need to move C1 left (C2 right)
		} else if L2 > R1 {
			hi = mid2 - 1 // A2's lower half too big; need to move C2 left.
		} else {
			return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2.0 // Otherwise, that's the right cut.
		}
	}
	return -1
}
