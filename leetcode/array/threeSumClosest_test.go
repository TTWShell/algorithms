package larray

import "testing"

func Test_threeSumClosest(t *testing.T) {
	input := []int{-1, 2, 1, -4}
	if r := threeSumClosest(input, 1); r != 2 {
		t.Fatal(r)
	}
}
