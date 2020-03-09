/* https://leetcode.com/problems/longest-consecutive-sequence/
Given an unsorted array of integers,
find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:

	Input: [100, 4, 200, 1, 3, 2]
	Output: 4
	Explanation: The longest consecutive elements sequence
	is [1, 2, 3, 4]. Therefore its length is 4.
*/

package larray

func longestConsecutive(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	cache := make(map[int]bool, length)
	for _, num := range nums {
		cache[num] = false
	}

	max := 1
	for _, num := range nums {
		if used := cache[num]; used {
			continue
		}
		cache[num] = true
		left, right := num, num
		for left = num - 1; ; left-- {
			if used, ok := cache[left]; ok && !used {
				cache[left] = true
			} else {
				left++
				break
			}
		}
		for right = num + 1; ; right++ {
			if used, ok := cache[right]; ok && !used {
				cache[right] = true
			} else {
				right--
				break
			}
		}
		if tmp := right - left + 1; tmp > max {
			max = tmp
		}
	}
	return max
}

// import (
// 	unionfind  "github.com/TTWShell/algorithms/data-structure/union-find"
// )

// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	uf := unionfind.New()

// 	sets := make(map[int]bool, len(nums))
// 	for _, num := range nums {
// 		sets[num] = true
// 		uf.Union(num, num)
// 	}

// 	for _, num := range nums {
// 		if _, ok := sets[num-1]; ok {
// 			uf.Union(num-1, num)
// 		}
// 		if _, ok := sets[num+1]; ok {
// 			uf.Union(num+1, num)
// 		}
// 	}

// 	res := 0
// 	for _, count := range uf.SetCount() {
// 		if count > res {
// 			res = count
// 		}
// 	}
// 	return res
// }
