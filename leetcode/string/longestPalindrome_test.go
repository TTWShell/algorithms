package lstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	input := []string{"a", "bb", "babad", "cbbd"}
	output := []string{"a", "bb", "bab", "bb"}

	for i := range input {
		if r := longestPalindrome(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
