package leetcode

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
}
