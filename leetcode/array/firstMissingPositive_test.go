package larray

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	if r := firstMissingPositive([]int{1, 2, 2, 1, 3, 1, 0, 4, 0}); r != 5 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{3, 4, 0, 2}); r != 1 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{-3, 9, 16, 4, 5, 16, -4, 9, 26, 2, 1, 19, -1, 25, 7, 22, 2, -7, 14, 2, 5, -6, 1, 17, 3, 24, -4, 17, 15}); r != 6 {
		t.Fatal(r)
	}
	if r := firstMissingPositive([]int{10, 4, 16, 54, 17, -7, 21, 15, 25, 31, 61, 1, 6, 12, 21, 46, 16, 56, 54, 12, 23, 20, 38, 63, 2, 27, 35, 11, 13, 47, 13, 11, 61, 39, 0, 14, 42, 8, 16, 54, 50, 12, -10, 43, 11, -1, 24, 38, -10, 13, 60, 0, 44, 11, 50, 33, 48, 20, 31, -4, 2, 54, -6, 51, 6}); r != 3 {
		t.Fatal(r)
	}
	if r := firstMissingPositive([]int{3, 1}); r != 2 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{2, 1}); r != 3 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{0, -1, 3, 1}); r != 2 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{1, 1}); r != 2 {
		t.Fatal(r)
	}
	if r := firstMissingPositive([]int{1, 2, 0}); r != 3 {
		t.Fatal(r)
	}

	if r := firstMissingPositive([]int{3, 4, -1, 1}); r != 2 {
		t.Fatal(r)
	}
}
