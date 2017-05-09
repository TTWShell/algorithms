/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

Subscribe to see which companies asked this question.
*/
package leetcode

func isValid(s string) bool {
	var str string
	for i := 0; i < len(s); i++ {
		if v := s[i]; v == '(' || v == '[' || v == '{' {
			str += string(v)
		} else if len(str) == 0 {
			return false
		} else {
			v1 := str[len(str)-1]
			if v == ')' && v1 != '(' {
				return false
			}
			if v == ']' && v1 != '[' {
				return false
			}
			if v == '}' && v1 != '{' {
				return false
			}
			str = str[:len(str)-1]
		}
	}
	return len(str) == 0
}
