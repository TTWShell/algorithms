/* https://leetcode.com/problems/insert-interval/
Given a set of non-overlapping intervals, insert a new interval into
the intervals (merge if necessary).

You may assume that the intervals were initially sorted according
to their start times.

Example 1:

	Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
	Output: [[1,5],[6,9]]

Example 2:

	Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]],
	newInterval = [4,8]
	Output: [[1,2],[3,10],[12,16]]
	Explanation: Because the new interval [4,8] overlaps
	with [3,5],[6,7],[8,10].

NOTE: input types have been changed on April 15, 2019.
Please reset to default code definition to get new method signature.
*/

package larray

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if len(intervals) == 0 || newInterval[1] < intervals[0][0] { // head
		res = append(res, newInterval)
	}

	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] { // 小端没有重叠
			res = append(res, intervals[i])
			// 需要直接插入newInterval的情况
			if i+1 < len(intervals) && newInterval[1] < intervals[i+1][0] {
				res = append(res, newInterval)
			}
			continue
		}
		if newInterval[1] < intervals[i][0] { // 大端没有重叠
			res = append(res, intervals[i])
			continue
		}

		// 处理重叠
		start := min(intervals[i][0], newInterval[0])
		i++
		for i+1 < len(intervals) &&
			(intervals[i+1][1] <= newInterval[1] ||
				newInterval[1] >= intervals[i+1][0]) {
			i++
		}
		i--
		if i+1 < len(intervals) && newInterval[1] >= intervals[i+1][0] {
			i++
		}
		end := max(intervals[i][1], newInterval[1])
		res = append(res, []int{start, end})
	}

	if res[len(res)-1][1] < newInterval[0] { // tail
		res = append(res, newInterval)
	}
	return res
}
