package leetcode

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	r := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	if len(r) != 3 || r[0] != -1 || r[1] != 3 || r[2] != -1 {
		t.Fatal(r)
	}

	r = nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
	if len(r) != 2 || r[0] != 3 || r[1] != -1 {
		t.Fatal(r)
	}

	r = nextGreaterElement([]int{2, 4}, []int{1, 4, 2, 3})
	if len(r) != 2 || r[0] != 3 || r[1] != -1 {
		t.Fatal(r)
	}
}
