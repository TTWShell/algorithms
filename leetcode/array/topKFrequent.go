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
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	countMaps := map[int]int{}

	for _, num := range nums {
		countMaps[num]++
	}

	h := &IntHeap{}
	heap.Init(h)
	for k, v := range countMaps {
		heap.Push(h, [2]int{k, v})
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		val := heap.Pop(h).([2]int)
		res[i] = val[0]
	}
	return res
}
