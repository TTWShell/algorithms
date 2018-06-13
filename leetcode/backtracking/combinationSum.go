/* https://leetcode.com/problems/combination-sum/description/
Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
    All numbers (including target) will be positive integers.
    The solution set must not contain duplicate combinations.

For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
    [
        [7],
        [2, 2, 3]
    ]
*/

package lbacktracking

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	for i := len(candidates) - 1; i >= 0; i-- {
		if now := candidates[i]; now > target {
			continue
		} else if now == target {
			res = append(res, []int{now})
		} else {
			results := combinationSum(candidates[0:i+1], target-now)
			if len(results) != 0 {
				for _, result := range results {
					res = append(res, append(result, now))
				}
			}
		}
	}
	return res
}
