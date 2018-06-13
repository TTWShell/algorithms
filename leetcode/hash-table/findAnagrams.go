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

package lht

func findAnagrams(s string, p string) []int {
	// https://discuss.leetcode.com/topic/64434/shortest-concise-java-o-n-sliding-window-solution
	r := []int{}
	if len(s) == 0 || len(p) == 0 {
		return r
	}
	// record each character in p to hash
	mp := make([]int, 26)
	for _, r := range p {
		mp[r-'a']++
	}

	// two points, initialize count to p's length
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		// move right everytime, if the character exists in p's hash, decrease the count
		// current hash value >= 1 means the character is existing in p
		if mp[s[right]-'a'] >= 1 {
			count--
		}
		mp[s[right]-'a']--
		right++

		// when the count is down to 0, means we found the right anagram
		// then add window's left to result list
		if count == 0 {
			r = append(r, left)
		}

		// if we find the window's size equals to p, then we have to move left (narrow the window) to find the new match window
		// ++ to reset the hash because we kicked out the left
		// only increase the count if the character is in p
		// the count >= 0 indicate it was original in the hash, cuz it won't go below 0
		if right-left == len(p) {
			if mp[s[left]-'a'] >= 0 {
				count++
			}
			mp[s[left]-'a']++
			left++
		}
	}
	return r
}

/* 400ms
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
*/
