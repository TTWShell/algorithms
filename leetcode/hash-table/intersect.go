/* https://leetcode.com/problems/intersection-of-two-arrays-ii/#/description
Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
    What if the given array is already sorted? How would you optimize your algorithm?
    What if nums1's size is small compared to nums2's size? Which algorithm is better?
    What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return []int{}
	}
	if m > n {
		m, n = n, len(nums1)
		temp := nums1
		nums1, nums2 = nums2, temp
	}
	map1, map2, r := make(map[int]int, m), make(map[int]int, n), make([]int, m)

	for i := range nums1 {
		map1[nums1[i]] += 1
	}

	for i := range nums2 {
		map2[nums2[i]] += 1
	}

	c := 0
	for k, v1 := range map1 {
		if v2, ok := map2[k]; ok {
			v := 0
			if v1 > v2 {
				v = v2
			} else {
				v = v1
			}

			for ; v > 0; v-- {
				r[c] = k
				c++
			}
		}
	}
	return r[:c]
}
