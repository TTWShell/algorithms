/* https://leetcode.com/problems/first-unique-character-in-a-string/#/description
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

    s = "leetcode"
    return 0.

    s = "loveleetcode",
    return 2.

Note: You may assume the string contain only lowercase letters.
*/

package leetcode

func firstUniqChar(s string) int {
	maps := make([]int, 26)
	for _, r := range s {
		maps[r-'a'] += 1
	}

	for i, r := range s {
		if maps[r-'a'] == 1 {
			return i
		}
	}
	return -1
}
