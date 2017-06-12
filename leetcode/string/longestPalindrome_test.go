package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	input := []string{"babad", "cbbd"}
	output := []string{"bab", "bb"}

	for i := range input {
		if r := longestPalindrome(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
