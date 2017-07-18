package leetcode

import "testing"

func Test_findMaxAverage(t *testing.T) {
	input := []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage(input, 4); r != 12.75 {
		t.Fatal(r)
	}
}
