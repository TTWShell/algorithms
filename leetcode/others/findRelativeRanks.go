/* https://leetcode.com/problems/relative-ranks/#/description
Given scores of N athletes, find their relative ranks and the people with the top three highest scores,
who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
    Input: [5, 4, 3, 2, 1]
    Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
    Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
    For the left two athletes, you just need to output their relative ranks according to their scores.

Note:
    N is a positive integer and won't exceed 10,000.
    All the scores of athletes are guaranteed to be unique.
*/

package lothers

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	n := len(nums)
	maps := make(map[int]int, n)
	for index, num := range nums {
		maps[num] = index
	}

	sort.Ints(nums)

	r := make([]string, n, n)
	for i, num := range nums {
		switch i {
		case n - 1:
			r[maps[num]] = "Gold Medal"
		case n - 2:
			r[maps[num]] = "Silver Medal"
		case n - 3:
			r[maps[num]] = "Bronze Medal"
		default:
			r[maps[num]] = strconv.Itoa(n - i)
		}
	}
	return r
}
