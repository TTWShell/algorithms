/* https://leetcode.com/problems/reverse-only-letters/description/
Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

Example 1:

	Input: "ab-cd"
	Output: "dc-ba"

Example 2:

	Input: "a-bC-dEf-ghIj"
	Output: "j-Ih-gfE-dCba"

Example 3:

	Input: "Test1ng-Leet=code-Q!"
	Output: "Qedo1ct-eeLg=ntse-T!"

Note:

	S.length <= 100
	33 <= S[i].ASCIIcode <= 122
	S doesn't contain \ or "
*/

package lstring

func reverseOnlyLetters(S string) string {
	l, r, res := 0, len(S)-1, make([]byte, len(S), len(S))
	for l <= r {
		if S[l] < 'A' || ('Z' < S[l] && S[l] < 'a') || S[l] > 'z' {
			res[l] = S[l]
			l++
			continue
		}
		if S[r] < 'A' || ('Z' < S[r] && S[r] < 'a') || S[r] > 'z' {
			res[r] = S[r]
			r--
			continue
		}
		res[l], res[r] = S[r], S[l]
		l++
		r--
	}
	return string(res)
}
