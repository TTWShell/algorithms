package leetcode

import "testing"

func Test_maxDistance(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
	output := 4

	if r := maxDistance(input); r != output {
		t.Fatal(r)
	}
}
