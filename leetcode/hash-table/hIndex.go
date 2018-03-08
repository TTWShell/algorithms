/* https://leetcode.com/problems/h-index/description/
Given an array of citations (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the definition of h-index on Wikipedia:
"A scientist has index h if h of his/her N papers have at least h citations each, and the other N − h papers have no more than h citations each."

For example, given citations = [3, 0, 6, 1, 5], which means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, his h-index is 3.

Note: If there are several possible values for h, the maximum one is taken as the h-index.
*/

package leetcode

/*
// 1. sort
import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := 0
	for i, citation := range citations {
		res = max(res, min(len(citations)-i, citation))
	}
	return res
}
*/

func hIndex(citations []int) int {
	count := make([]int, len(citations)+1)
	for _, citation := range citations {
		if citation >= len(citations) {
			count[len(citations)]++ // 当引用数大于等于 n 时，我们均将其数量计入 count[n] 中
		} else {
			count[citation]++
		}
	}
	for i := len(citations); i > 0; i-- {
		if count[i] >= i {
			return i
		}
		count[i-1] += count[i] // 引用数大于 i-1 的数量是i-1及之后的累加
	}
	return 0
}
