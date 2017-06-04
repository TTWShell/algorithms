package leetcode

import "testing"

func Test_arrangeCoins(t *testing.T) {
	input := []int{5, 8}
	output := []int{2, 3}

	for i := range input {
		if r := arrangeCoins(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
