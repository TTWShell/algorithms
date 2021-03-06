/* https://leetcode.com/problems/maximum-length-of-pair-chain/#/description
You are given n pairs of numbers. In every pair, the first number is always smaller than the second number.

Now, we define a pair (c, d) can follow another pair (a, b) if and only if b < c. Chain of pairs can be formed in this fashion.

Given a set of pairs, find the length longest chain which can be formed. You needn't use up all the given pairs. You can select pairs in any order.

Example 1:
    Input: [[1,2], [2,3], [3,4]]
    Output: 2
    Explanation: The longest chain is [1,2] -> [3,4]
Note:
    The number of given pairs will be in the range [1, 1000].
*/

package larray

import (
	"math"
	"sort"
)

type SortPair [][]int

func (s SortPair) Len() int {
	return len(s)
}

func (s SortPair) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s SortPair) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func findLongestChain(pairs [][]int) int {
	// sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	sort.Sort(SortPair(pairs))

	last, res := []int{math.MinInt32, math.MinInt32}, 0
	for _, p := range pairs {
		if p[0] > last[1] {
			res++
			last = p
		}
	}
	return res
}
