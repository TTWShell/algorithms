package leetcode

import "testing"

func Test_fourSum(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	output := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}

	r := fourSum(input, 0)

	for i, row := range output {
		for j := range row {
			if output[i][j] != r[i][j] {
				t.Fatal(r)
			}
		}
	}
}
