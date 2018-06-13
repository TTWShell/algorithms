/* https://leetcode.com/problems/combination-sum-ii/description/
Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

Each number in C may only be used once in the combination.

Note:
    All numbers (including target) will be positive integers.
    The solution set must not contain duplicate combinations.

For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
    [
      [1, 7],
      [1, 2, 5],
      [2, 6],
      [1, 1, 6]
    ]
*/

package lbacktracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	for i := len(candidates) - 1; i >= 0; i-- {
		if now := candidates[i]; now > target {
			continue
		} else if now == target {
			res = append(res, []int{now})
		} else {
			results := combinationSum2(candidates[0:i], target-now)
			if len(results) != 0 {
				for _, result := range results {
					res = append(res, append(result, now))
				}
			}
		}
		for i-1 >= 0 && candidates[i] == candidates[i-1] {
			i--
		}
	}
	return res
}
