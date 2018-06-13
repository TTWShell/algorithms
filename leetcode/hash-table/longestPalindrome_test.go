package lht

import "testing"

func Test_longestPalindrome(t *testing.T) {
	input := []string{"abccccdd", "aaabaaaa", "AAA", "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"}
	result := []int{7, 7, 3, 55}

	for i := range input {
		if r := longestPalindrome(input[i]); r != result[i] {
			t.Fatal(input[i], result[i], r)
		}
	}
}
