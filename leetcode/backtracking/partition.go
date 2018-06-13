/* https://leetcode.com/problems/palindrome-partitioning/description/
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

For example, given s = "aab",
Return

[
  ["aa","b"],
  ["a","a","b"]
]
*/

package lbacktracking

func partition(s string) [][]string {
	res := [][]string{}
	for i := len(s); i >= 1; i-- {
		if tmpS := s[0:i]; partitionHelper(tmpS) {
			if i == len(s) {
				res = append(res, []string{tmpS})
				continue
			}
			if subs := partition(s[i:]); len(subs) != 0 {
				for _, sub := range subs {
					res = append(res, append([]string{tmpS}, sub...))
				}
			}
		}
	}
	return res
}

func partitionHelper(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
