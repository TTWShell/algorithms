/* https://leetcode.com/problems/ransom-note/#/description
Given an arbitrary ransom note string and another string containing letters from all the magazines,
write a functioransom-noten that will return true if the ransom note can be constructed from the magazines ;
otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
    You may assume that both strings contain only lowercase letters.

    canConstruct("a", "b") -> false
    canConstruct("aa", "ab") -> false
    canConstruct("aa", "aab") -> true
*/

package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	maps := make([]int, 26)
	for _, r := range magazine {
		maps[r-'a'] += 1
	}

	for _, r := range ransomNote {
		if v := maps[r-'a']; v > 0 {
			maps[r-'a'] -= 1
		} else {
			return false
		}
	}
	return true
}
