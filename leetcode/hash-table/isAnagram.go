/* https://leetcode.com/problems/valid-anagram/#/description
Given two strings s and t, write a function to determine if t is an anagram of s.

For example,
s = "anagram", t = "nagaram", return true.
s = "rat", t = "car", return false.

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

package leetcode

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n {
		return false
	}

	maps := make(map[byte]int)
	for i := range s {
		maps[s[i]]++
		maps[t[i]]--
	}

	for _, v := range maps {
		if v != 0 {
			return false
		}
	}
	return true
}
