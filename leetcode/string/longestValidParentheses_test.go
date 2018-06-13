package lstring

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	if r := longestValidParentheses("()"); r != 2 {
		t.Fatal(r)
	}

	if r := longestValidParentheses("(()(((()"); r != 2 {
		t.Fatal(r)
	}

	if r := longestValidParentheses("(()"); r != 2 {
		t.Fatal(r)
	}

	if r := longestValidParentheses(")()())"); r != 4 {
		t.Fatal(r)
	}

	if r := longestValidParentheses("()(()"); r != 2 {
		t.Fatal(r)
	}
}
