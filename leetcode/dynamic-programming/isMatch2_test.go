package leetcode

import "testing"

func Test_isMatch2(t *testing.T) {
	input := [][]string{
		[]string{"aa", "a"},
		[]string{"aa", "aa"},
		[]string{"aaa", "aa"},
		[]string{"aa", "*"},
		[]string{"aa", "a*"},
		[]string{"aa", "?*"},
		[]string{"aab", "c*a*b"},
	}
	output := []bool{false, true, false, true, true, true, false}

	for i := range input {
		if r := isMatch2(input[i][0], input[i][1]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
