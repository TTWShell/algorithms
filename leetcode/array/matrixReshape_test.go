package larray

import "testing"

func Test_matrixReshape(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}

	r := matrixReshape(nums, 2, 4)
	for i := range nums {
		for j := range nums[i] {
			if nums[i][j] != r[i][j] {
				t.Fatal(r)
			}
		}
	}

	r = matrixReshape(nums, 1, 4)
	output := []int{1, 2, 3, 4}
	for i := range output {
		if output[i] != r[0][i] {
			t.Fatal(r)
		}
	}

	r = matrixReshape([][]int{{1, 2, 3, 4}}, 2, 2)
	for i := range nums {
		for j := range nums[i] {
			if nums[i][j] != r[i][j] {
				t.Fatal(r)
			}
		}
	}
}
