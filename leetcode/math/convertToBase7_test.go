package leetcode

import "testing"

func Test_convertToBase7(t *testing.T) {
	input := []int{100, -7, 0}
	output := []string{"202", "-10", "0"}

	for i := range input {
		if r := convertToBase7(input[i]); r != output[i] {
			t.Fatal(input[i], r, output[i])
		}
	}
}
