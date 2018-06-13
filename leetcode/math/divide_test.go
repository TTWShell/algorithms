package lmath

import "testing"

func Test_divide(t *testing.T) {
	if r := divide(32, 3); r != 10 {
		t.Fatal(r)
	}
}
