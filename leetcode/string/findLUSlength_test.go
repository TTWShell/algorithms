package leetcode

import "testing"

func Test_findLUSlength(t *testing.T) {
	if r := findLUSlength("aba", "cdc"); r != 3 {
		t.Fatal(r)
	}
}
