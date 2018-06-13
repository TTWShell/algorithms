/* https://leetcode.com/problems/letter-case-permutation/description/
Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.  Return a list of all possible strings we could create.

Examples:

    Input: S = "a1b2"
    Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

    Input: S = "3z4"
    Output: ["3z4", "3Z4"]

    Input: S = "12345"
    Output: ["12345"]

Note:

    S will be a string with length at most 12.
    S will consist only of letters or digits.
*/

package lstring

func letterCasePermutation(S string) []string {
	length := len(S)
	if length == 0 {
		return []string{""}
	}

	res := []string{}
	lastLetter := S[length-1]
	tmp := []string{string(lastLetter)}

	switch {
	case 97 <= lastLetter && lastLetter <= 122:
		tmp = append(tmp, string(lastLetter-32))
	case 65 <= lastLetter && lastLetter <= 90:
		tmp = append(tmp, string(lastLetter+32))
	}

	if len(S) == 1 {
		return tmp
	}

	for _, pre := range letterCasePermutation(S[:length-1]) {
		for _, cur := range tmp {
			res = append(res, pre+cur)
		}
	}
	return res
}
