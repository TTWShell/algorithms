/* https://leetcode.com/problems/kth-largest-element-in-an-array/description/
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
For example,
Given [3,2,1,5,6,4] and k = 2, return 5.

Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

package leetcode

import "github.com/TTWShell/algorithms/heap"

type findKthLargestInt int

func (i findKthLargestInt) LessThan(e heap.Element) bool {
	return i < e.(findKthLargestInt)
}

func findKthLargest(nums []int, k int) int {
	heap := heap.New(true) // min heap

	for _, number := range nums {
		num := findKthLargestInt(number)
		if heap.Len() < k {
			heap.Insert(num)
			continue
		}
		if heap.Peek().(findKthLargestInt) < num {
			heap.Extract()
			heap.Insert(num)
		}
	}
	return int(heap.Peek().(findKthLargestInt))
}
