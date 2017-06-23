package leetcode

import "testing"

func Test_maxArea(t *testing.T) {
	height := [][]int{{1, 2}, {1, 2, 3, 4, 5, 3, 2, 1}}
	output := []int{1, 10}

	for i := range height {
		if r := maxArea(height[i]); r != output[i] {
			t.Fatal(r)
		}
	}
}
