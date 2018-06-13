package lstring

import "testing"

func Test_canConstruct(t *testing.T) {
	input := [][]string{{"a", "b"}, {"aa", "ab"}, {"aa", "aab"}}
	result := []bool{false, false, true}

	for i := range input {
		if r := canConstruct(input[i][0], input[i][1]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
