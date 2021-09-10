/* https://leetcode.com/problems/find-common-characters/description/
Given an array A of strings made only from lowercase letters,
return a list of all characters that show up in all strings within the list (including duplicates).
For example, if a character occurs 3 times in all strings but not 4 times,
you need to include that character three times in the final answer.

You may return the answer in any order.

Example 1:

	Input: ["bella","label","roller"]
	Output: ["e","l","l"]

Example 2:

	Input: ["cool","lock","cook"]
	Output: ["c","o"]

Note:

	1 <= A.length <= 100
	1 <= A[i].length <= 100
	A[i][j] is a lowercase letter
*/
package lht

func commonChars(A []string) []string {
	cache := [26]int{}
	for idx, word := range A {
		tmp := [26]int{}
		for i := range word {
			tmp[word[i]-97]++
		}
		if idx == 0 {
			cache = tmp
		} else {
			for i := 0; i < 26; i++ {
				if tmp[i] < cache[i] {
					cache[i] = tmp[i]
				}
			}
		}
	}

	res := []string{}
	for k, count := range cache {
		for i := 0; i < count; i++ {
			res = append(res, string(rune(k+97)))
		}
	}
	return res
}
