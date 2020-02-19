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
	pq "github.com/TTWShell/algorithms/data-structure/priority-queue"
)

func topKFrequent(nums []int, k int) []int {
	countMaps := map[int]int{}

	for _, num := range nums {
		countMaps[num]++
	}

	pQueue := pq.PQConstructor()
	for k, v := range countMaps {
		pQueue.Insert(pq.Element{Value: k, Priority: v})
	}

	res := make([]int, k)
	for i := range res {
		res[i] = pQueue.Extract().Value.(int)
	}
	return res
}
