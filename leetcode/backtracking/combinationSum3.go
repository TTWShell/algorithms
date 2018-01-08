/* https://leetcode.com/problems/combination-sum-iii/description/
Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.

Example 1:
	Input: k = 3, n = 7
	Output:
	[[1,2,4]]

Example 2:
	Input: k = 3, n = 9
	Output:
	[[1,2,6], [1,3,5], [2,3,4]]
*/

package leetcode

func combinationSum3(k int, n int) [][]int {
	var helper func(k, n, start int) [][]int
	helper = func(k, n, start int) [][]int {
		if k == 1 {
			if start <= n && n <= 9 {
				return [][]int{{n}}
			}
			return [][]int{}
		}

		res := [][]int{}
		for i := start; i <= 9; i++ {
			if n-i < 1 {
				break
			}
			for _, subSum := range helper(k-1, n-i, i+1) {
				res = append(res, append([]int{i}, subSum...))
			}
		}
		return res
	}

	return helper(k, n, 1)
}
