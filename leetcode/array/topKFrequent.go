/* https://leetcode.com/problems/top-k-frequent-elements/
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

Example 2:

	Input: nums = [1], k = 1
	Output: [1]

Note:

	You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
	Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

package larray

import (
	"sort"
)

type matrixSlice [][2]int

func (m matrixSlice) Len() int {
	return len(m)
}
func (m matrixSlice) Less(i, j int) bool {
	return m[i][1] > m[j][1]
}

func (m matrixSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func topKFrequent(nums []int, k int) []int {
	countMaps := map[int]int{}

	for _, num := range nums {
		countMaps[num]++
	}

	ms := make(matrixSlice, len(countMaps))
	idx := 0
	for k, v := range countMaps {
		ms[idx] = [2]int{k, v}
		idx++
	}

	sort.Sort(ms)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = ms[i][0]
	}
	return res
}
