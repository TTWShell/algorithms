package larray

import "testing"

func Test_maxDistance(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
	output := 4

	if r := maxDistance(input); r != output {
		t.Fatal(r)
	}

	input = [][]int{{1, 4}, {0, 5}}
	if r := maxDistance(input); r != 4 {
		t.Fatal(r)
	}
}
