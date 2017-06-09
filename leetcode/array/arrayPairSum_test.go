package leetcode

import "testing"

func Test_arrayPairSum(t *testing.T) {
	input := [][]int{{1, 4, 3, 2}}
	output := []int{4}

	for i := range input {
		if r := arrayPairSum(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
