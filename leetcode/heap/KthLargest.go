/* https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream.
For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

Example:

    int k = 3;
    int[] arr = [4,5,8,2];
    KthLargest kthLargest = new KthLargest(3, arr);
    kthLargest.add(3);   // returns 4
    kthLargest.add(5);   // returns 5
    kthLargest.add(10);  // returns 5
    kthLargest.add(9);   // returns 8
    kthLargest.add(4);   // returns 8

Note:
    You may assume that nums' length ≥ k-1 and k ≥ 1.
*/

package lheap

import "sort"

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	if len(nums) > k {
		nums = nums[len(nums)-k:] // only save last k
	}
	return KthLargest{k: k, nums: nums}
}

func (this *KthLargest) Add(val int) int {
	idx := this.findIdx(val)
	// idx := sort.Search(len(this.nums), func(i int) bool { return this.nums[i] > val })

	if idx == 0 && len(this.nums) < this.k {
		this.nums = append([]int{val}, this.nums...)
	}

	if idx > 0 {
		this.nums = append(this.nums, 0)
		copy(this.nums[idx+1:], this.nums[idx:])
		this.nums[idx] = val
		this.nums = this.nums[len(this.nums)-this.k:]
	}

	return this.nums[0]
}

func (this *KthLargest) findIdx(val int) int {
	// find first element > val, suppport duplicate element
	if len(this.nums) == 0 {
		return 0
	}

	start, end := 0, len(this.nums)-1
	if val < this.nums[start] {
		return start
	}
	if val > this.nums[end] {
		return end + 1
	}

	for end > start {
		mid := start + (end-start)>>1
		if this.nums[mid] > val {
			end = mid - 1
		} else if this.nums[mid] == val {
			idx := mid
			for ; idx < len(this.nums) && this.nums[idx] == val; idx++ {
			}
			return idx
		} else {
			start = mid + 1
		}
	}

	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		if this.nums[i] > val {
			return i
		}
	}
	return end + 1
}
