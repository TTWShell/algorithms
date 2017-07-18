package leetcode

import "testing"

func Test_findMaxAverage2(t *testing.T) {
	input := []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 4); r != 12.75 {
		t.Fatal(r)
	}

	input = []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 5); r != 10.8 {
		t.Fatal(r)
	}

	input = []int{1, 12, -5, -6, 50, 3}
	if r := findMaxAverage2(input, 6); r != 9.166666666666666 {
		t.Fatal(r)
	}

	input = []int{8, 9, 3, 1, 8, 3, 0, 6, 9, 2}
	if r := findMaxAverage2(input, 8); r != 5.222222222222222 {
		t.Fatal(r)
	}
}
