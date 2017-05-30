package leetcode

import "testing"

func Test_findNthDigit(t *testing.T) {
	input := []int{3, 11}
	result := []int{3, 0}

	for i := range input {
		if r := findNthDigit(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
