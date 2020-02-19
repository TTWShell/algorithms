/* https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
Find the length of the longest substring T of a given string (consists of lowercase letters only)
such that every character in T appears no less than k times.

Example 1:

	Input:
	s = "aaabb", k = 3

	Output:
	3

	The longest substring is "aaa", as 'a' is repeated 3 times.

Example 2:

	Input:
	s = "ababbc", k = 2

	Output:
	5

	The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
*/

package lstring

func longestSubstring(s string, k int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var dfs func(s string, left, right int, k int) int
	dfs = func(s string, left, right int, k int) int {
		letters := [26]int{}
		for i := left; i < right; i++ {
			letters[s[i]-97]++
		}

		valid := true
		for i := 0; i < 26; i++ {
			if letters[i] > 0 && letters[i] < k {
				valid = false
				break
			}
		}
		if valid {
			return right - left
		}

		var length int
		start := left
		for i := left; i < right; i++ {
			if letters[s[i]-97] < k {
				length = max(length, dfs(s, start, i, k))
				start = i + 1
			}
		}
		return max(length, dfs(s, start, right, k))
	}

	return dfs(s, 0, len(s), k)
}

// Runtime: 336 ms, faster than 12.82%, Memory Usage: 2 MB, less than 100.00%
// func longestSubstring(s string, k int) int {
// 	tmp := [26]int{}
// 	for i := range s {
// 		tmp[s[i]-97]++
// 	}

// 	for i := len(s); i >= k; i-- {
// 		for idx := i; idx < len(s); idx++ { // 去掉cache中不需要的tail
// 			tmp[s[idx]-97]--
// 		}
// 		subEnd := len(s) - i
// 		for j := 0; j <= subEnd; j++ {
// 			if j != 0 { // 去掉不需要的cache
// 				tmp[s[j-1]-97]--
// 			}

// 			flag := true
// 			for idx := range tmp {
// 				if tmp[idx] != 0 && tmp[idx] < k {
// 					flag = false
// 					break
// 				}
// 			}
// 			if last := j + i; 0 < last && last < len(s) { // 增加下一步的cache
// 				tmp[s[last]-97]++
// 			}
// 			if flag {
// 				return i
// 			}
// 		}

// 		for idx := 0; idx < len(s)-i; idx++ { // 补全head
// 			tmp[s[idx]-97]++
// 		}
// 	}
// 	return 0
// }
