/* https://leetcode.com/problems/maximum-product-subarray/description/
Find the contiguous subarray within an array (containing at least one number) which has the largest product.

For example, given the array [2,3,-2,4],
the contiguous subarray [2,3] has the largest product = 6.
*/

package leetcode

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	maxMul, minMul, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		a, b := maxMul*nums[i], minMul*nums[i]
		maxMul = max(max(a, b), nums[i])
		minMul = min(min(a, b), nums[i])
		res = max(res, maxMul)
	}
	return res
}
