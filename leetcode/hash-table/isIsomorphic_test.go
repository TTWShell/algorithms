package leetcode

import "testing"

func Test_isIsomorphic(t *testing.T) {
	input := [][]string{[]string{"aba", "baa"}, []string{"aa", "ab"}, []string{"egg", "add"}, []string{"foo", "bar"}, []string{"paper", "title"}}
	result := []bool{false, false, true, false, true}

	for i := 0; i < len(input); i++ {
		if r := isIsomorphic(input[i][0], input[i][1]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
