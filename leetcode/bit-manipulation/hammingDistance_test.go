package leetcode

import "testing"

func Test_hammingDistance(t *testing.T) {
	if r := hammingDistance(1, 4); r != 2 {
		t.Fatal(r)
	}
}
