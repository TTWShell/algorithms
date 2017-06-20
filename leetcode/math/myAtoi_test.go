package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	input := []string{"  123", "  +-2", "123-+", "123dad234324", "2147483648"}
	output := []int{123, 0, 123, 123, 2147483647}

	for i := range input {
		if r := myAtoi(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
