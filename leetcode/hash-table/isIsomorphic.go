/* https://leetcode.com/problems/isomorphic-strings/#/description
Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

For example,
Given "egg", "add", return true.

Given "foo", "bar", return false.

Given "paper", "title", return true.

Note:
You may assume both s and t have the same length.
*/

package leetcode

func isIsomorphic(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}

	mmap, nmap := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		if _, ok := mmap[s[i]]; !ok {
			mmap[s[i]] = t[i]
		}
		if _, ok := nmap[t[i]]; !ok {
			nmap[t[i]] = s[i]
		}

		if t[i] != mmap[s[i]] {
			return false
		}
		if s[i] != nmap[t[i]] {
			return false
		}
	}
	return true
}
