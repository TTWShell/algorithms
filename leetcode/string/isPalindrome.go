/* https://leetcode.com/problems/valid-palindrome/#/description
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.

Note:
Have you consider that the string might be empty? This is a good question to ask during an interview.

For the purpose of this problem, we define empty string as valid palindrome.
*/

package lstring

func isPalindrome(s string) bool {
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		if !isAlpha(bytes[i]) {
			i++
			continue
		}
		if !isAlpha(bytes[j]) {
			j--
			continue
		}
		if toLower(bytes[i]) != toLower(bytes[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}

func isAlpha(b byte) bool {
	if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}
