/* https://leetcode.com/problems/longest-harmonious-subsequence/#/description
We define a harmonious array is an array where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:
    Input: [1,3,2,2,5,2,3,7]
    Output: 5
    Explanation: The longest harmonious subsequence is [3,2,2,2,3].

Note: The length of the input array will not exceed 20,000.
*/

package lht

func findLHS(nums []int) int {
	maps := make(map[int]int, len(nums))
	for _, num := range nums {
		maps[num]++
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	result := 0
	for k, v := range maps {
		if count, ok := maps[k+1]; ok {
			result = max(result, v+count)
		}
	}
	return result
}
