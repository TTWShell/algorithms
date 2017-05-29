/* https://leetcode.com/problems/range-sum-query-immutable/#/description
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:
    Given nums = [-2, 0, 3, -5, 2, -1]

    sumRange(0, 2) -> 1
    sumRange(2, 5) -> -1
    sumRange(0, 5) -> -3

Note:
    You may assume that the array does not change.
    There are many calls to sumRange function.
*/

package leetcode

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	a := NumArray{}
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	a.sums = nums
	return a
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > j || i > len(this.sums)-1 {
		panic("i cannot > j or i index outof range")
	}

	var m, n int
	if i <= 0 {
		m = 0
	} else {
		m = this.sums[i-1]
	}
	if j > len(this.sums)-1 {
		n = this.sums[len(this.sums)-1]
	} else {
		n = this.sums[j]
	}

	return n - m
}
