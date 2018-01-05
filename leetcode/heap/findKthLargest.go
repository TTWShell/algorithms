/* https://leetcode.com/problems/kth-largest-element-in-an-array/description/
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
For example,
Given [3,2,1,5,6,4] and k = 2, return 5.

Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

package leetcode

// 快速选择算法
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// 以第一个数为基准，大的放在左边
	idx := 0
	for i := 1; i < len(nums); i++ {
		if next := idx + 1; nums[i] > nums[idx] && next < len(nums) {
			nums[next], nums[i] = nums[i], nums[next]
			nums[idx], nums[next] = nums[next], nums[idx]
			idx++
		}
	}

	if target := k - 1; idx == target {
		return nums[idx]
	} else if idx > target {
		return findKthLargest(nums[:idx], k)
	}
	return findKthLargest(nums[idx+1:], k-idx-1)
}

/*
// 基于堆的实现
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
*/
