package lht

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	input := [][]int{{0, 0}, {1, 0}, {2, 0}}
	if r := numberOfBoomerangs(input); r != 2 {
		t.Fatal(input, r)
	}
}
