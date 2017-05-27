package leetcode

import "testing"

func Test_trailingZeroes(t *testing.T) {
	input := []int{1, 5, 10, 15, 25, 100}
	result := []int{0, 1, 2, 3, 6, 24}

	for i := 0; i < len(input); i++ {
		if r := trailingZeroes(input[i]); r != result[i] {
			t.Fatal(input[i], result[i], r)
		}
	}
}
