package lstring

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	input := []string{"ffffF", "Google", "g", "leetcode", "USA", "FlaG"}
	output := []bool{false, true, true, true, true, false}

	for i := range input {
		if r := detectCapitalUse(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
