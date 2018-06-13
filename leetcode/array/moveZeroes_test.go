package larray

import "testing"

func Test_moveZeroes(t *testing.T) {
	nums := [][]int{{0, 1, 0, 3, 12}, {0, 0, 1}}
	result := [][]int{{1, 3, 12, 0, 0}, {1, 0, 0}}

	for i := range nums {
		moveZeroes(nums[i])
		for j := range nums[i] {
			if nums[i][j] != result[i][j] {
				t.Fatal(nums[i], result[i])
			}
		}
	}
}
