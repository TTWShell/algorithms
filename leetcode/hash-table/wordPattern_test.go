package lht

import "testing"

func Test_wordPattern(t *testing.T) {
	input := [][]string{
		{"abba", "dog cat cat dog"},
		{"abba", "dog cat cat fish"},
		{"aaaa", "dog cat cat dog"},
		{"abba", "dog dog dog dog"},
	}
	result := []bool{true, false, false, false}

	for i := range input {
		if r := wordPattern(input[i][0], input[i][1]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
