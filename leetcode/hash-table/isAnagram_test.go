package lht

import "testing"

func Test_isAnagram(t *testing.T) {
	input := [][]string{{"anagram", "nagaram"}, {"rat", "car"}}
	result := []bool{true, false}

	for i := range input {
		if r := isAnagram(input[i][0], input[i][1]); r != result[i] {
			t.Fatal(input, r)
		}
	}
}
