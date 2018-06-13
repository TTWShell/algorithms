package lmath

import "testing"

func Test_maxCount(t *testing.T) {
	if r := maxCount(3, 3, [][]int{{2, 2}, {3, 3}}); r != 4 {
		t.Fatal(r)
	}
}
