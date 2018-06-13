package larray

import "testing"

func Test_trap(t *testing.T) {
	if r := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); r != 6 {
		t.Fatal(r)
	}
}
