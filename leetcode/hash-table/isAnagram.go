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

package lht

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make([]int, 26)
	count := len(s)
	for _, v := range s {
		record[v-'a']++ // -'a' 后变为0-25的数字
	}
	for _, v := range t {
		if record[v-'a'] >= 1 {
			count--
		}
		record[v-'a']--
	}
	return count == 0
}
