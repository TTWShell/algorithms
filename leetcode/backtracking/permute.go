/* https://leetcode.com/problems/permutations/description/
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]
*/

package lbacktracking

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	res := [][]int{}
	for i, num := range nums {
		tmp := []int{num}
		args := append([]int{}, nums[0:i]...)
		for _, r := range permute(append(args, nums[i+1:]...)) {
			res = append(res, append(tmp, r...))
		}
	}
	return res
}
