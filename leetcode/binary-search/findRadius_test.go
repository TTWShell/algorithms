package leetcode

import "testing"

func Test_findRadius(t *testing.T) {
	if r := findRadius([]int{1, 2, 3}, []int{2}); r != 1 {
		t.Fatal(r)
	}
	if r := findRadius([]int{1, 2, 3, 4}, []int{1, 4}); r != 1 {
		t.Fatal(r)
	}

}
