/* https://leetcode.com/problems/combinations/description/
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:

[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

package leetcode

func combine(n int, k int) [][]int {
	if n < 1 || k < 1 || k > n {
		return [][]int{}
	}

	var helper func(start, end, k int) [][]int
	helper = func(start, end, k int) [][]int {
		res := [][]int{}
		if k == 1 {
			for i := start; i <= end; i++ {
				res = append(res, []int{i})
			}
		} else {
			for i := start; i <= start+(end-start+1)-k; i++ {
				tmp := []int{i}
				for _, sub := range helper(i+1, end, k-1) {
					res = append(res, append(tmp, sub...))
				}
			}
		}
		return res
	}

	return helper(1, n, k)
}
