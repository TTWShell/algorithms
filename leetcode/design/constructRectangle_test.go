package ldesign

import "testing"

func Test_constructRectangle(t *testing.T) {
	if r := constructRectangle(4); len(r) != 2 || r[0] != 2 || r[1] != 2 {
		t.Fatal(4, r)
	}
}
