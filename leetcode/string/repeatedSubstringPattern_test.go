package leetcode

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	input := []string{"abaababaab", "abacababacab", "abac", "abab", "aba", "abcabcabcabc"}
	output := []bool{true, true, false, true, false, true}

	for i := range input {
		if r := repeatedSubstringPattern(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
