package leetcode

import "testing"

func Test_countSegments(t *testing.T) {
	input := []string{"Hello, my name is John "}
	output := []int{5}

	for i := range input {
		if r := countSegments(input[i]); r != output[i] {
			t.Fatal(input[i], r)
		}
	}
}
