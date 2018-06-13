/* https://leetcode.com/problems/intersection-of-two-arrays/#/description
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.
*/

package lht

func intersection(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return []int{}
	}
	if m > n {
		m, n = n, len(nums1)
		temp := nums1
		nums1, nums2 = nums2, temp
	}
	maps, r := make(map[int]int, m), make([]int, m)

	for i := range nums1 {
		if _, ok := maps[nums1[i]]; !ok {
			maps[nums1[i]] = 1
		}
	}

	c := 0
	for i := range nums2 {
		if v, ok := maps[nums2[i]]; ok {
			maps[nums2[i]] = v + 1
			if v+1 == 2 {
				r[c] = nums2[i]
				c++
			}
		}
	}

	return r[:c]
}
