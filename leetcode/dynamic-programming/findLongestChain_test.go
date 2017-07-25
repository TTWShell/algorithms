package leetcode

import "testing"

func Test_findLongestChain(t *testing.T) {
	if r := findLongestChain([][]int{{1, 2}, {2, 3}, {3, 4}}); r != 2 {
		t.Fatal(r)
	}
}
