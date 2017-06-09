package leetcode

import "testing"

func Test_findPairs(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1}, {3, 1, 4, 1, 5}, {1, 2, 3, 4, 5}, {1, 3, 1, 5, 4}}
	ks := []int{0, 2, 1, 0}
	output := []int{1, 2, 4, 1}

	for i := range input {
		if r := findPairs(input[i], ks[i]); r != output[i] {
			t.Fatal(input[i], ks[i], output[i], r)
		}
	}
}
