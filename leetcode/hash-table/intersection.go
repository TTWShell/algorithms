/* https://leetcode.com/problems/intersection-of-two-arrays/#/description
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

Note:
Each element in the result must be unique.
The result can be in any order.
*/

package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	maps, r := make(map[int]int, len(nums1)+len(nums2)), []int{}

	for i := range nums1 {
		if _, ok := maps[nums1[i]]; !ok {
			maps[nums1[i]] = 1
		}
	}
	for i := range nums2 {
		if v, ok := maps[nums2[i]]; ok {
			maps[nums2[i]] = v + 1
		}
	}

	for k, v := range maps {
		if v > 1 {
			r = append(r, k)
		}
	}

	return r
}
