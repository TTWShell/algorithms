package leetcode

import "testing"

func Test_findContentChildren(t *testing.T) {
	input := [][]int{{1, 2, 3}, {1, 1}}
	if r := findContentChildren(input[0], input[1]); r != 1 {
		t.Fatal(input, r)
	}

	input = [][]int{{1, 2}, {1, 2, 3}}
	if r := findContentChildren(input[0], input[1]); r != 2 {
		t.Fatal(input, r)
	}
}
