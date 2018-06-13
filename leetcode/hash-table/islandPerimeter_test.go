package lht

import "testing"

func Test_islandPerimeter(t *testing.T) {
	input := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	if r := islandPerimeter(input); r != 16 {
		t.Fatal(r)
	}
}
