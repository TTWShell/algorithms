package leetcode

import "testing"
import "strings"

func Test_generateParenthesis(t *testing.T) {
	n := 3
	output := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	r := generateParenthesis(n)
	if len(r) != len(output) || strings.Join(r, "") != strings.Join(output, "") {
		t.Fatal(r)
	}
}
