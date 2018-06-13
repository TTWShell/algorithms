package larray

import "testing"

func Test_checkPossibility(t *testing.T) {
	if r := checkPossibility([]int{4, 2, 3}); r != true {
		t.Fatal(r)
	}

	if r := checkPossibility([]int{4, 2, 1}); r != false {
		t.Fatal(r)
	}

	if r := checkPossibility([]int{3, 4, 2, 3}); r != false {
		t.Fatal(r)
	}
}
