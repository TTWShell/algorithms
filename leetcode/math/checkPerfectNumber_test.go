package lmath

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	input := []int{1, 28}
	output := []bool{false, true}

	for i := range input {
		if r := checkPerfectNumber(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
