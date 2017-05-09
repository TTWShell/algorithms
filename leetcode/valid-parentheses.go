/*  https://leetcode.com/problems/valid-parentheses/#/description
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/
package leetcode

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		v1 := stack[len(stack)-1]
		if (v == ')' && v1 != '(') || (v == ']' && v1 != '[') || (v == '}' && v1 != '{') {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
