package ldp

import "testing"

func Test_isMatch(t *testing.T) {
	input := [][]string{
		[]string{"aa", "a"},
		[]string{"aa", "aa"},
		[]string{"aaa", "aa"},
		[]string{"aa", "a*"},
		[]string{"aa", ".*"},
		[]string{"ab", ".*"},
		[]string{"aab", "c*a*b"},
	}
	output := []bool{false, true, false, true, true, true, true}

	for i := range input {
		if r := isMatch(input[i][0], input[i][1]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
