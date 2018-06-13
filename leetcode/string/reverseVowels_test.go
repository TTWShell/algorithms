package lstring

import "testing"

func Test_reverseVowels(t *testing.T) {
	input := []string{"hello", "leetcode", "aA"}
	result := []string{"holle", "leotcede", "Aa"}

	for i := 0; i < len(input); i++ {
		if r := reverseVowels(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
