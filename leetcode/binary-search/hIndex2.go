/* https://leetcode.com/problems/h-index-ii/description/
Follow up for H-Index: What if the citations array is sorted in ascending order? Could you optimize your algorithm?
*/

package leetcode

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

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

	res := 0
	for l, r := 0, len(citations)-1; l <= r; {
		mid := l + (r-l)>>1
		res = max(res, min(len(citations)-mid, citations[mid]))
		if len(citations)-mid > citations[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
