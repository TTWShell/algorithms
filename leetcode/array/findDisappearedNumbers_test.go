package leetcode

import "testing"

func Test_findDisappearedNumbers(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	output := []int{5, 6}

	r := findDisappearedNumbers(input)

	if len(r) != len(output) {
		t.Fatal(r, output)
	}

	for i := range output {
		if r[i] != output[i] {
			t.Fatal(r, output)
		}
	}
}
