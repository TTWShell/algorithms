/* https://leetcode.com/problems/find-all-anagrams-in-a-string/#/description
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:
    Input:
    s: "cbaebabacd" p: "abc"
    Output:
    [0, 6]

Explanation:
    The substring with start index = 0 is "cba", which is an anagram of "abc".
    The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
    Input:
    s: "abab" p: "ab"
    Output:
    [0, 1, 2]

Explanation:
    The substring with start index = 0 is "ab", which is an anagram of "ab".
    The substring with start index = 1 is "ba", which is an anagram of "ab".
    The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

package leetcode

func findAnagrams(s string, p string) []int {
	mp := make([]int, 26)
	for _, r := range p {
		mp[r-'a']++
	}

	r := []int{}
	for i := 0; i < len(s)-len(p)+1; i++ {
		temp := make([]int, 26)
		j := i
		for ; j < i+len(p); j++ {
			t := s[j] - 'a'
			temp[t]++
			if mp[t] == 0 {
				i = j
				break
			}
			if mp[t] < temp[t] {
				break
			}
		}
		if j == i+len(p) {
			r = append(r, i)
		}
	}
	return r
}
