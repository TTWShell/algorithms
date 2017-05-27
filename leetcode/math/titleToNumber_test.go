package leetcode

import "testing"

func Test_titleToNumber(t *testing.T) {
	input := []string{"A", "B", "C", "Z", "AA", "AB"}
	result := []int{1, 2, 3, 26, 27, 28}

	for i := 0; i < len(input); i++ {
		if r := titleToNumber(input[i]); r != result[i] {
			t.Fatal(input[i], result[i], r)
		}
	}
}
