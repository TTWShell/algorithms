package leetcode

import "testing"

func Test_thirdMax(t *testing.T) {
	input := [][]int{{3, 2, 1}, {1, 2}, {2, 2, 3, 1}}
	output := []int{1, 2, 1}

	for i := range input {
		if r := thirdMax(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
