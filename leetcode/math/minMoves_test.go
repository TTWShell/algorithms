package leetcode

import "testing"

func Test_minMoves(t *testing.T) {
	input := []int{1, 2, 3}
	output := 3
	if r := minMoves(input); r != output {
		t.Fatal(input, r)
	}
}
