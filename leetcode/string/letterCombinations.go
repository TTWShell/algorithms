/* https://leetcode.com/problems/letter-combinations-of-a-phone-number/#/description
Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.

http://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png

    Input:Digit string "23"
    Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
    Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package leetcode

func letterCombinations(digits string) []string {
	maps := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	res := []string{}
	for i, digit := range digits {
		if i == 0 {
			for _, letter := range maps[digit-'0'] {
				res = append(res, string(letter))
			}
		} else {
			temp := []string{}
			for _, agoLetter := range res {
				for _, letter := range maps[digit-'0'] {
					temp = append(temp, agoLetter+string(letter))
				}
			}
			res = temp
		}
	}
	return res
}
