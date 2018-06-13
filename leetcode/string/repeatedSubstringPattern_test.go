package lstring

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	input := []string{"ababac", "abaababaab", "abacababacab", "abac", "abab", "aba", "abcabcabcabc"}
	output := []bool{false, true, true, false, true, false, true}

	for i := range input {
		if r := repeatedSubstringPattern(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
