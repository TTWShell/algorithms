package leetcode

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	if r := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}); r != 3 {
		t.Fatal(r)
	}
}
