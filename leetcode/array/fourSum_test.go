package larray

import "testing"

func Test_fourSum(t *testing.T) {
	test := func(input []int, output [][]int) {
		r := fourSum(input, 0)
		for i, row := range output {
			for j := range row {
				if output[i][j] != r[i][j] {
					t.Fatal(r)
				}
			}
		}
	}

	input := []int{1, 0, -1, 0, -2, 2}
	output := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	test(input, output)

	input = []int{-3, -2, -1, 0, 0, 1, 2, 3}
	output = [][]int{
		{-3, -2, 2, 3},
		{-3, -1, 1, 3},
		{-3, 0, 0, 3},
		{-3, 0, 1, 2},
		{-2, -1, 0, 3},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	test(input, output)
}
