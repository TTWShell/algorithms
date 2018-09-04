/* https://leetcode.com/problems/range-sum-query-mutable/description/
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

	Given nums = [1, 3, 5]

	sumRange(0, 2) -> 9
	update(1, 2)
	sumRange(0, 2) -> 8
Note:

	The array is only modifiable by the update function.
	You may assume the number of calls to update and sumRange function is distributed evenly.
*/

package lbit

func lowbit(x int) int {
	return x & (-x)
}

type NumArray struct {
	c      []int
	length int
}

func Constructor(nums []int) NumArray {
	length := len(nums) + 1
	c := make([]int, length, length)
	for i := 1; i < length; i++ {
		c[i] = nums[i-1]
		for j := i - 2; j >= i-lowbit(i); j-- {
			c[i] += nums[j]
		}
	}
	return NumArray{c: c, length: length}
}

func (this *NumArray) Update(i int, val int) {
	delta := val - (this.SumRange(i, i))
	for j := i + 1; j < this.length; j += lowbit(j) {
		this.c[j] += delta
		if lowbit(j) == 0 {
			break
		}
	}
}

func (this *NumArray) Sum(i int) int {
	res := 0
	for j := i; j > 0; j -= lowbit(j) {
		res += this.c[j]
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Sum(j+1) - this.Sum(i)
}
