package leetcode

import "testing"

func Test_canCross(t *testing.T) {
	input := [][]int{{0, 1, 3, 5, 6, 8, 12, 17}, {0, 1, 2, 3, 4, 8, 9, 11}}
	result := []bool{true, false}

	for i := range input {
		if r := canCross(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
