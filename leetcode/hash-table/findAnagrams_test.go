package lht

import "testing"

func Test_findAnagrams(t *testing.T) {
	input := [][]string{{"cbaebabacd", "abc"}, {"abab", "ab"}}
	output := [][]int{{0, 6}, {0, 1, 2}}

	for i := range input {
		r := findAnagrams(input[i][0], input[i][1])
		if len(r) != len(output[i]) {
			t.Fatal(r, output[i])
		}
		for j := range r {
			if r[j] != output[i][j] {
				t.Fatal(r, output[i])
			}
		}
	}
}
