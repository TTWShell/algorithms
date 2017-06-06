package leetcode

import "testing"

func Test_findComplement(t *testing.T) {
	input := []int{5, 1}
	output := []int{2, 1}

	for i := range input {
		if r := findComplement(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
