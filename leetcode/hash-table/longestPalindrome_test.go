package leetcode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	input := []string{"abccccdd", "aaabaaaa"}
	result := []int{7, 7}

	for i := range input {
		if r := longestPalindrome(input[i]); r != result[i] {
			t.Fatal(input[i], result[i], r)
		}
	}
}
