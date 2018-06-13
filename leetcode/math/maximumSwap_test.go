package lmath

import "testing"

func Test_maximumSwap(t *testing.T) {
	if r := maximumSwap(2736); r != 7236 {
		t.Fatal(r)
	}

	if r := maximumSwap(9973); r != 9973 {
		t.Fatal(r)
	}
}
