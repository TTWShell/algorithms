package lstring

import "testing"

func Test_letterCombinations(t *testing.T) {
	output := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	r := letterCombinations("23")

	for i := range output {
		if output[i] != r[i] {
			t.Fatal(r)
		}
	}

}
