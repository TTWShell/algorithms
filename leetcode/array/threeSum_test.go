package larray

import "testing"

func Test_threeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	output := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	r := threeSum(input)
	for i, row := range output {
		for j := range row {
			if r[i][j] != row[j] {
				t.Fatal(r)
			}
		}
	}

	input = []int{0, 0, 0, 0}
	output = [][]int{{0, 0, 0}}
	if r := threeSum(input); len(r) != 1 {
		t.Fatal(r)
	}

	input = []int{-1, 0, 1, 0}
	output = [][]int{{-1, 0, 1}}
	if r := threeSum(input); len(r) != 1 {
		t.Fatal(r)
	}

	input = []int{3, 0, -2, -1, 1, 2}
	output = [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}
	r = threeSum(input)
	for i, row := range output {
		for j := range row {
			if r[i][j] != row[j] {
				t.Fatal(r)
			}
		}
	}
}
