package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	if r := firstMissingPositive([]int{1, 2, 0}); r != 3 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{3, 4, -1, 1}); r != 2 {
		t.Fatal(r)
	}
}
