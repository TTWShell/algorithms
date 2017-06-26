package leetcode

import "testing"

func Test_maximumProduct(t *testing.T) {
	input := [][]int{{1, 2, 3}, {1, 2, 3, 4}, {-4, -3, -2, -1, 60}}
	output := []int{6, 24, 720}

	for i, nums := range input {
		if r := maximumProduct(nums); r != output[i] {
			t.Fatal(nums, r)
		}
	}
}
