/* https://leetcode.com/problems/merge-intervals/description/
Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
*/

package larray

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

import "sort"

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (intervals Intervals) Len() int {
	return len(intervals)
}

func (intervals Intervals) Less(i, j int) bool {
	return intervals[i].Start < intervals[j].Start
}

func (intervals Intervals) Swap(i, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func merge2(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(Intervals(intervals))

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	res := []Interval{}
	var tmp *Interval
	for _, interval := range intervals {
		if tmp == nil {
			tmp = &Interval{interval.Start, interval.End}
			continue
		}
		if interval.Start > tmp.End {
			res = append(res, *tmp)
			tmp = &Interval{interval.Start, interval.End}
		} else {
			tmp.Start = min(tmp.Start, interval.Start)
			tmp.End = max(tmp.End, interval.End)
		}
	}
	res = append(res, *tmp)
	return res
}
