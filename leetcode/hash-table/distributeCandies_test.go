package leetcode

import "testing"

func Test_distributeCandies(t *testing.T) {
	input := [][]int{{1, 1, 2, 2, 3, 3}, {1, 1, 2, 3}}
	output := []int{3, 2}

	for i := range input {
		if r := distributeCandies(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
