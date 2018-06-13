package lmath

import "testing"

func Test_findNthDigit(t *testing.T) {
	input := []int{3, 10, 11, 12, 13, 14, 15, 1000}
	result := []int{3, 1, 0, 1, 1, 1, 2, 3}

	for i := range input {
		if r := findNthDigit(input[i]); r != result[i] {
			t.Fatal(input[i], r)
		}
	}
}
